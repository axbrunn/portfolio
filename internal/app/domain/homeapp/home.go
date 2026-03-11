package homeapp

import (
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/ui/html/pages"
)

type handlers struct {
	log *slog.Logger
}

func (h *handlers) home(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(r.Context(), w)
}
