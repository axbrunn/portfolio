package blogrepo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/axbrunn/portfolio/internal/domain"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]domain.Post, error) {
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

	var posts []domain.Post
	for rows.Next() {
		var p domain.Post
		err := rows.Scan(&p.ID, &p.Title, &p.Slug, &p.Excerpt, &p.Body, &p.Published, &p.CreatedAt, &p.UpdatedAt, &p.PublishedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (r *Repository) GetBySlug(ctx context.Context, slug string) (domain.Post, error) {
	stmt := `
		SELECT id, title, slug, excerpt, body, published, created_at, updated_at, published_at
		FROM posts
		WHERE slug = ?`

	row := r.db.QueryRowContext(ctx, stmt, slug)

	var p domain.Post
	err := row.Scan(
		&p.ID, &p.Title, &p.Slug, &p.Excerpt, &p.Body, &p.Published, &p.CreatedAt, &p.UpdatedAt, &p.PublishedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Post{}, domain.ErrNoRecord
		} else {
			return domain.Post{}, err
		}
	}
	return p, err
}

func (r *Repository) Insert(ctx context.Context, p domain.Post) (string, error) {
	stmt := `INSERT INTO posts (title, slug, excerpt, body, published, published_at) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, stmt, p.Title, p.Slug, p.Excerpt, p.Body, p.Published, p.PublishedAt)
	if err != nil {
		return "", err
	}

	return p.Slug, nil
}

func (r *Repository) Update(ctx context.Context, p domain.Post) (string, error) {
	stmt := `
		UPDATE posts
		SET title = ?, excerpt = ?, body = ?, published = ?, published_at = ?, updated_at = CURRENT_TIMESTAMP
		WHERE slug = ?`

	_, err := r.db.ExecContext(ctx, stmt, p.Title, p.Excerpt, p.Body, p.Published, p.PublishedAt, p.Slug)
	if err != nil {
		return "", err
	}

	return p.Slug, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	stmt := `DELETE FROM posts WHERE slug = ?`

	_, err := r.db.ExecContext(ctx, stmt, id)
	return err
}
