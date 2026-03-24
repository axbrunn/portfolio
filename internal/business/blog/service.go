package blog

import (
	"context"
	"errors"
	"time"

	"github.com/axbrunn/portfolio/internal/business/validator"
)

type BlogService struct {
	repo Repository
}

func New(repo Repository) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) ViewAll(ctx context.Context) ([]BlogPost, error) {
	return s.repo.SelectAll(ctx)
}

func (s *BlogService) View(ctx context.Context, slug string) (BlogPost, error) {
	return s.repo.SelectBySlug(ctx, slug)
}

func (s *BlogService) CreatePost(ctx context.Context, title, slug, excerpt, body string, published bool) (string, error) {
	v := &validator.Validator{}

	v.CheckField(validator.NotBlank(title), "title", "titel is verplicht")
	v.CheckField(validator.NotBlank(slug), "slug", "slug is verplicht")
	v.CheckField(validator.NotBlank(excerpt), "excerpt", "excerpt is verplicht")
	v.CheckField(validator.NotBlank(body), "body", "body is verplicht")

	// Controleer of de slug al in gebruik is.
	_, err := s.repo.SelectBySlug(ctx, slug)
	if err == nil {
		v.AddFieldError("slug", "slug is al in gebruik")
	} else if !errors.Is(err, ErrNoRecord) {
		return "", err
	}

	if !v.Valid() {
		return "", &ValidationError{Fields: v.FieldErrors}
	}

	p := BlogPost{
		Title:     title,
		Slug:      slug,
		Excerpt:   excerpt,
		Body:      body,
		Published: published,
	}

	// Business regel: alleen een gepubliceerde post krijgt een publicatiedatum.
	if p.Published {
		now := time.Now().UTC()
		p.PublishedAt = &now
	}

	return s.repo.Insert(ctx, p)
}

func (s *BlogService) Update(ctx context.Context, id uint) (BlogPost, error) {
	return s.repo.SelectByID(ctx, id)
}

func (s *BlogService) UpdatePut(ctx context.Context, p BlogPost) (string, error) {
	return s.repo.Update(ctx, p)
}

func (s *BlogService) Delete(ctx context.Context, id uint) error {
	return s.repo.DeleteByID(ctx, id)
}
