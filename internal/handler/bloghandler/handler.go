package bloghandler

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/axbrunn/portfolio/internal/domain"
	"github.com/axbrunn/portfolio/ui/html/pages"
)

type BlogService interface {
	GetAll(ctx context.Context) ([]domain.Post, error)
	GetBySlug(ctx context.Context, slug string) (domain.Post, error)
	Create(ctx context.Context, p domain.Post) (domain.Post, error)
	Update(ctx context.Context, p domain.Post) error
	Delete(ctx context.Context, slug string) error
}

// postResponse is the handler's own type for rendering.
// Keeps template concerns out of domain.Post.
type postResponse struct {
	Title       string
	Slug        string
	Excerpt     string
	Body        string
	PublishedAt *time.Time
}

func toResponse(p domain.Post) postResponse {
	return postResponse{
		Title:       p.Title,
		Slug:        p.Slug,
		Excerpt:     p.Excerpt,
		Body:        p.Body,
		PublishedAt: p.PublishedAt,
	}
}

// createPostRequest is the handler's own type for incoming create requests.
type createPostRequest struct {
	Title   string
	Slug    string
	Excerpt string
	Body    string
}

func (req createPostRequest) toDomain() domain.Post {
	return domain.Post{
		Title:   req.Title,
		Slug:    req.Slug,
		Excerpt: req.Excerpt,
		Body:    req.Body,
	}
}

// updatePostRequest is the handler's own type for incoming update requests.
type updatePostRequest struct {
	Title     string
	Excerpt   string
	Body      string
	Published bool
}

func (req updatePostRequest) toDomain(slug string) domain.Post {
	return domain.Post{
		Slug:      slug,
		Title:     req.Title,
		Excerpt:   req.Excerpt,
		Body:      req.Body,
		Published: req.Published,
	}
}

// =============================================================================

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

	pages.Blog(toResponse(post).Title).Render(r.Context(), w)
}

func (h *handlers) create(w http.ResponseWriter, r *http.Request) {
}

func (h *handlers) update(w http.ResponseWriter, r *http.Request) {
}

func (h *handlers) delete(w http.ResponseWriter, r *http.Request) {
}
