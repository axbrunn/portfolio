package web

import (
	"net/http"

	"github.com/axbrunn/portfolio/internal/app"
	"github.com/axbrunn/portfolio/internal/web/handlers"
	"github.com/axbrunn/portfolio/ui"
)

func Routes(app *app.Application) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /", handlers.Home)

	return mux
}
