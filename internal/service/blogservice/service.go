package blogservice

import (
	"context"

	"github.com/axbrunn/portfolio/internal/domain"
)

type PostRepository interface {
	GetAll(ctx context.Context) ([]domain.Post, error)
	GetBySlug(ctx context.Context, slug string) (domain.Post, error)
	Create(ctx context.Context, p domain.Post) (domain.Post, error)
	Update(ctx context.Context, p domain.Post) error
	Delete(ctx context.Context, slug string) error
}

type Service struct {
	repo PostRepository
}

func New(repo PostRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Post, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetBySlug(ctx context.Context, slug string) (domain.Post, error) {
	return s.repo.GetBySlug(ctx, slug)
}

func (s *Service) Create(ctx context.Context, p domain.Post) (domain.Post, error) {
	return s.repo.Create(ctx, p)
}

func (s *Service) Update(ctx context.Context, p domain.Post) error {
	return s.repo.Update(ctx, p)
}

func (s *Service) Delete(ctx context.Context, slug string) error {
	return s.repo.Delete(ctx, slug)
}
