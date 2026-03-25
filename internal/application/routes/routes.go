package routes

import (
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/internal/application/handlers"
	"github.com/axbrunn/portfolio/internal/application/middleware"
	"github.com/axbrunn/portfolio/ui"
)

type Router struct {
	Logger      *slog.Logger
	HomeHandler *handlers.Home
	BlogHandler *handlers.Blog
}

func New(r Router) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /{$}", r.HomeHandler.Home)

	mux.HandleFunc("GET /blog/{$}", r.BlogHandler.ViewAll)
	mux.HandleFunc("GET /blog/{slug}", r.BlogHandler.View)
	mux.HandleFunc("GET /blog/create/", r.BlogHandler.Create)
	mux.HandleFunc("POST /blog/create", r.BlogHandler.CreatePost)
	mux.HandleFunc("GET /blog/update/{id}", r.BlogHandler.Update)
	mux.HandleFunc("PUT /blog/update/{id}", r.BlogHandler.UpdatePut)
	mux.HandleFunc("DELETE /blog/{id}", r.BlogHandler.Delete)

	return middleware.Logger(r.Logger, mux)
}
