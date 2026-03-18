package blog

import (
	"context"

	"github.com/axbrunn/portfolio/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.BlogPost, error)
	GetBySlug(ctx context.Context, slug string) (domain.BlogPost, error)
	Post(ctx context.Context, p domain.BlogPost) (string, error)
	Put(ctx context.Context, p domain.BlogPost) (string, error)
	Delete(ctx context.Context, slug string) error
}

type Repository interface {
	SelectAll(ctx context.Context) ([]domain.BlogPost, error)
	SelectBySlug(ctx context.Context, slug string) (domain.BlogPost, error)
	Insert(ctx context.Context, p domain.BlogPost) (string, error)
	Update(ctx context.Context, p domain.BlogPost) (string, error)
	Delete(ctx context.Context, slug string) error
}
