package blog

import (
	"context"
	"database/sql"
	"errors"

	"github.com/axbrunn/portfolio/internal/domain"
)

type BlogRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) SelectAll(ctx context.Context) ([]domain.BlogPost, error) {
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

	var posts []domain.BlogPost
	for rows.Next() {
		var p domain.BlogPost
		err := rows.Scan(&p.ID, &p.Title, &p.Slug, &p.Excerpt, &p.Body, &p.Published, &p.CreatedAt, &p.UpdatedAt, &p.PublishedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (r *BlogRepository) SelectBySlug(ctx context.Context, slug string) (domain.BlogPost, error) {
	stmt := `
		SELECT id, title, slug, excerpt, body, published, created_at, updated_at, published_at
		FROM posts
		WHERE slug = ?`

	row := r.db.QueryRowContext(ctx, stmt, slug)

	var p domain.BlogPost
	err := row.Scan(
		&p.ID, &p.Title, &p.Slug, &p.Excerpt, &p.Body, &p.Published, &p.CreatedAt, &p.UpdatedAt, &p.PublishedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.BlogPost{}, domain.ErrNoRecord
		} else {
			return domain.BlogPost{}, err
		}
	}
	return p, err
}

func (r *BlogRepository) Insert(ctx context.Context, p domain.BlogPost) (string, error) {
	stmt := `INSERT INTO posts (title, slug, excerpt, body, published, published_at)
	VALUES (?, ?, ?, ?, ?, UTC_TIMESTAMP())`

	_, err := r.db.ExecContext(ctx, stmt, p.Title, p.Slug, p.Excerpt, p.Body, p.Published, p.PublishedAt)
	if err != nil {
		return "", err
	}

	return p.Slug, nil
}

func (r *BlogRepository) Update(ctx context.Context, p domain.BlogPost) (string, error) {
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

func (r *BlogRepository) Delete(ctx context.Context, id string) error {
	stmt := `DELETE FROM posts WHERE slug = ?`

	_, err := r.db.ExecContext(ctx, stmt, id)
	return err
}
