package blog

import (
	"context"
	"database/sql"
	"errors"

	bizblog "github.com/axbrunn/portfolio/internal/business/blog"
)

type BlogRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) SelectAll(ctx context.Context) ([]bizblog.BlogPost, error) {
	stmt := `
		SELECT id, title, slug, excerpt, body, published, created_at, updated_at, published_at
		FROM posts
		WHERE published = true
		ORDER BY published_at DESC`

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []bizblog.BlogPost
	for rows.Next() {
		// Scan in het store-interne model zodat DB-types (sql.NullTime) correct worden gelezen.
		var m blogPostModel
		err := rows.Scan(&m.ID, &m.Title, &m.Slug, &m.Excerpt, &m.Body, &m.Published, &m.CreatedAt, &m.UpdatedAt, &m.PublishedAt)
		if err != nil {
			return nil, err
		}
		// Converteer naar business type voordat we het teruggeven aan de service.
		posts = append(posts, toBusiness(m))
	}

	return posts, nil
}

func (r *BlogRepository) SelectBySlug(ctx context.Context, slug string) (bizblog.BlogPost, error) {
	stmt := `
		SELECT id, title, slug, excerpt, body, published, created_at, updated_at, published_at
		FROM posts
		WHERE slug = ?`

	row := r.db.QueryRowContext(ctx, stmt, slug)

	// Scan in het store-interne model zodat DB-types (sql.NullTime) correct worden gelezen.
	var m blogPostModel
	err := row.Scan(&m.ID, &m.Title, &m.Slug, &m.Excerpt, &m.Body, &m.Published, &m.CreatedAt, &m.UpdatedAt, &m.PublishedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return bizblog.BlogPost{}, bizblog.ErrNoRecord
		}
		return bizblog.BlogPost{}, err
	}

	// Converteer naar business type voordat we het teruggeven aan de service.
	return toBusiness(m), nil
}

func (r *BlogRepository) SelectByID(ctx context.Context, id uint) (bizblog.BlogPost, error) {
	stmt := `
		SELECT id, title, slug, excerpt, body, published, created_at, updated_at, published_at
		FROM posts
		WHERE id = ?`

	row := r.db.QueryRowContext(ctx, stmt, id)

	// Scan in het store-interne model zodat DB-types (sql.NullTime) correct worden gelezen.
	var m blogPostModel
	err := row.Scan(&m.ID, &m.Title, &m.Slug, &m.Excerpt, &m.Body, &m.Published, &m.CreatedAt, &m.UpdatedAt, &m.PublishedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return bizblog.BlogPost{}, bizblog.ErrNoRecord
		}
		return bizblog.BlogPost{}, err
	}

	// Converteer naar business type voordat we het teruggeven aan de service.
	return toBusiness(m), nil
}

func (r *BlogRepository) Insert(ctx context.Context, p bizblog.BlogPost) (string, error) {
	// Converteer naar model zodat PublishedAt als sql.NullTime wordt meegegeven.
	// Als de post niet gepubliceerd is, heeft de business layer PublishedAt op nil gelaten
	// en schrijft de mapper dat als NULL naar de DB.
	m := toModel(p)

	stmt := `INSERT INTO posts (title, slug, excerpt, body, published, published_at)
	VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, stmt, m.Title, m.Slug, m.Excerpt, m.Body, m.Published, m.PublishedAt)
	if err != nil {
		return "", err
	}

	return p.Slug, nil
}

func (r *BlogRepository) Update(ctx context.Context, p bizblog.BlogPost) (string, error) {
	// Converteer business type naar model zodat we sql.NullTime kunnen meegeven aan de query.
	m := toModel(p)

	stmt := `
		UPDATE posts
		SET title = ?, excerpt = ?, body = ?, published = ?, published_at = ?, updated_at = CURRENT_TIMESTAMP
		WHERE slug = ?`

	_, err := r.db.ExecContext(ctx, stmt, m.Title, m.Excerpt, m.Body, m.Published, m.PublishedAt, m.Slug)
	if err != nil {
		return "", err
	}

	return p.Slug, nil
}

func (r *BlogRepository) DeleteByID(ctx context.Context, id uint) error {
	stmt := `DELETE FROM posts WHERE id = ?`

	_, err := r.db.ExecContext(ctx, stmt, id)
	return err
}
