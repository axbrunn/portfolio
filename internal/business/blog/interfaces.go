package blog

import (
	"context"
)

type Service interface {
	GetAll(ctx context.Context) ([]BlogPost, error)
	GetBySlug(ctx context.Context, slug string) (BlogPost, error)
	Post(ctx context.Context, p BlogPost) (string, error)
	Put(ctx context.Context, p BlogPost) (string, error)
	Delete(ctx context.Context, slug string) error
}

type Repository interface {
	SelectAll(ctx context.Context) ([]BlogPost, error)
	SelectBySlug(ctx context.Context, slug string) (BlogPost, error)
	Insert(ctx context.Context, p BlogPost) (string, error)
	Update(ctx context.Context, p BlogPost) (string, error)
	Delete(ctx context.Context, slug string) error
}
