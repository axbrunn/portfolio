package bloghandler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/internal/domain"
	"github.com/axbrunn/portfolio/ui/html/pages"
)

type BlogService interface {
	GetAll(ctx context.Context) ([]domain.Post, error)
	GetBySlug(ctx context.Context, slug string) (domain.Post, error)
	Insert(ctx context.Context, p domain.Post) (string, error)
	Update(ctx context.Context, p domain.Post) (string, error)
	Delete(ctx context.Context, slug string) error
}

type handlers struct {
	log     *slog.Logger
	service BlogService
}

func (h *handlers) getAll(w http.ResponseWriter, r *http.Request) {
}

func (h *handlers) get(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	post, err := h.service.GetBySlug(r.Context(), slug)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	pages.Blog(post.Title).Render(r.Context(), w)
}

func (h *handlers) insert(w http.ResponseWriter, r *http.Request) {
}

func (h *handlers) update(w http.ResponseWriter, r *http.Request) {
}

func (h *handlers) delete(w http.ResponseWriter, r *http.Request) {
}
