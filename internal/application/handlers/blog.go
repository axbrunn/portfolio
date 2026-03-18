package handlers

import (
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/internal/business/blog"
	"github.com/axbrunn/portfolio/ui/html/pages"
)

type Blog struct {
	log *slog.Logger
	svc blog.Service
}

func NewBlog(log *slog.Logger, svc blog.Service) *Blog {
	return &Blog{log: log, svc: svc}
}

func (h *Blog) GetAll(w http.ResponseWriter, r *http.Request) {
}

func (h *Blog) Get(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	post, err := h.svc.GetBySlug(r.Context(), slug)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	pages.Blog(post.Title).Render(r.Context(), w)
}

func (h *Blog) Insert(w http.ResponseWriter, r *http.Request) {
}

func (h *Blog) Update(w http.ResponseWriter, r *http.Request) {
}

func (h *Blog) Delete(w http.ResponseWriter, r *http.Request) {
}
