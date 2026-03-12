package blogapp

import (
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/ui/html/pages"
)

type handlers struct {
	log *slog.Logger
}

func (h *handlers) getAll(w http.ResponseWriter, r *http.Request) {
}

func (h *handlers) get(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	pages.Blog(slug).Render(r.Context(), w)
}

func (h *handlers) create(w http.ResponseWriter, r *http.Request) {
}

func (h *handlers) update(w http.ResponseWriter, r *http.Request) {
}

func (h *handlers) delete(w http.ResponseWriter, r *http.Request) {
}
