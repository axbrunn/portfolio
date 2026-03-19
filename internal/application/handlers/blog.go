package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/internal/application/helpers"
	"github.com/axbrunn/portfolio/internal/business/blog"
	"github.com/axbrunn/portfolio/internal/domain"
	"github.com/axbrunn/portfolio/ui/html/pages"
)

type Blog struct {
	logger *slog.Logger
	svc    blog.Service
}

func NewBlog(log *slog.Logger, svc blog.Service) *Blog {
	return &Blog{
		logger: log,
		svc:    svc,
	}
}

func (h *Blog) ViewAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.svc.GetAll(r.Context())
	if err != nil {
		helpers.ServerError(h.logger, w, r, err)
		return
	}

	for _, post := range posts {
		fmt.Fprintf(w, "%+v\n", post)
	}

}

func (h *Blog) View(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	post, err := h.svc.GetBySlug(r.Context(), slug)
	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			helpers.ServerError(h.logger, w, r, err)
		}
		return
	}

	pages.Blog(post.Title).Render(r.Context(), w)
}

func (h *Blog) Create(w http.ResponseWriter, r *http.Request) {
}

func (h *Blog) CreatePost(w http.ResponseWriter, r *http.Request) {
}

func (h *Blog) Update(w http.ResponseWriter, r *http.Request) {
}

func (h *Blog) Delete(w http.ResponseWriter, r *http.Request) {
}
