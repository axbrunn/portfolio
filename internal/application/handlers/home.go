package handlers

import (
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/ui/html/pages"
)

type Home struct {
	log *slog.Logger
}

func NewHome(log *slog.Logger) *Home {
	return &Home{log: log}
}

func (h *Home) Home(w http.ResponseWriter, r *http.Request) {
	pages.HomeView().Render(r.Context(), w)
}
