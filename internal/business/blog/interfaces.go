package blog

import (
	"context"
)

type Service interface {
	ViewAll(ctx context.Context) ([]BlogPost, error)
	View(ctx context.Context, slug string) (BlogPost, error)
	CreatePost(ctx context.Context, title, slug, excerpt, body string, published bool) (string, error)
	Update(ctx context.Context, id uint) (BlogPost, error)
	UpdatePut(ctx context.Context, p BlogPost) (string, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	SelectAll(ctx context.Context) ([]BlogPost, error)
	SelectBySlug(ctx context.Context, slug string) (BlogPost, error)
	SelectByID(ctx context.Context, id uint) (BlogPost, error)
	Insert(ctx context.Context, p BlogPost) (string, error)
	Update(ctx context.Context, p BlogPost) (string, error)
	DeleteByID(ctx context.Context, id uint) error
}
