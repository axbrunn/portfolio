package blog

import (
	"context"
)

type BlogService struct {
	repo Repository
}

func New(repo Repository) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) GetAll(ctx context.Context) ([]BlogPost, error) {
	return s.repo.SelectAll(ctx)
}

func (s *BlogService) GetBySlug(ctx context.Context, slug string) (BlogPost, error) {
	return s.repo.SelectBySlug(ctx, slug)
}

func (s *BlogService) Post(ctx context.Context, p BlogPost) (string, error) {
	return s.repo.Insert(ctx, p)
}

func (s *BlogService) Put(ctx context.Context, p BlogPost) (string, error) {
	return s.repo.Update(ctx, p)
}

func (s *BlogService) Delete(ctx context.Context, slug string) error {
	return s.repo.Delete(ctx, slug)
}
