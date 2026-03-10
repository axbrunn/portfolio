package handlers

import (
	"net/http"

	"github.com/axbrunn/portfolio/ui/html/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(r.Context(), w)
}
