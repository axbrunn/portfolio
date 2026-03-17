package blogrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/axbrunn/portfolio/internal/domain"
)

// dbPost is the repository's own type for scanning SQL rows.
// It handles nullable DB types so domain.Post stays clean.
type dbPost struct {
	ID          uint
	Title       string
	Slug        string
	Excerpt     string
	Body        string
	Published   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt sql.NullTime
}

func (p dbPost) toDomain() domain.Post {
	post := domain.Post{
		ID:        p.ID,
		Title:     p.Title,
		Slug:      p.Slug,
		Excerpt:   p.Excerpt,
		Body:      p.Body,
		Published: p.Published,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	if p.PublishedAt.Valid {
		post.PublishedAt = &p.PublishedAt.Time
	}
	return post
}

// =============================================================================

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]domain.Post, error) {
	query := `
		SELECT id, title, slug, excerpt, body, published, created_at, updated_at, published_at
		FROM posts
		WHERE published = true
		ORDER BY published_at DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var p dbPost
		err := rows.Scan(&p.ID, &p.Title, &p.Slug, &p.Excerpt, &p.Body, &p.Published, &p.CreatedAt, &p.UpdatedAt, &p.PublishedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p.toDomain())
	}

	return posts, rows.Err()
}

func (r *Repository) GetBySlug(ctx context.Context, slug string) (domain.Post, error) {
	query := `
		SELECT id, title, slug, excerpt, body, published, created_at, updated_at, published_at
		FROM posts
		WHERE slug = ?`

	var p dbPost
	err := r.db.QueryRowContext(ctx, query, slug).Scan(
		&p.ID, &p.Title, &p.Slug, &p.Excerpt, &p.Body, &p.Published, &p.CreatedAt, &p.UpdatedAt, &p.PublishedAt,
	)
	return p.toDomain(), err
}

func (r *Repository) Create(ctx context.Context, p domain.Post) (domain.Post, error) {
	query := `INSERT INTO posts (title, slug, excerpt, body, published, published_at) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, p.Title, p.Slug, p.Excerpt, p.Body, p.Published, p.PublishedAt)
	if err != nil {
		return domain.Post{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Post{}, err
	}

	p.ID = uint(id)
	return p, nil
}

func (r *Repository) Update(ctx context.Context, p domain.Post) error {
	query := `
		UPDATE posts
		SET title = ?, excerpt = ?, body = ?, published = ?, published_at = ?, updated_at = CURRENT_TIMESTAMP
		WHERE slug = ?`

	_, err := r.db.ExecContext(ctx, query, p.Title, p.Excerpt, p.Body, p.Published, p.PublishedAt, p.Slug)
	return err
}

func (r *Repository) Delete(ctx context.Context, slug string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM posts WHERE slug = ?`, slug)
	return err
}
