package routes

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/internal/application/handlers"
	"github.com/axbrunn/portfolio/internal/application/middleware"
	blogbus "github.com/axbrunn/portfolio/internal/business/blog"
	blogrepo "github.com/axbrunn/portfolio/internal/store/blog"
	"github.com/axbrunn/portfolio/ui"
)

type Config struct {
	Log *slog.Logger
	DB  *sql.DB
}

func New(cfg Config) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	blogrepo := blogrepo.New(cfg.DB)
	blogsvc := blogbus.New(blogrepo)

	homeHandler := handlers.NewHome(cfg.Log)
	blogHandler := handlers.NewBlog(cfg.Log, blogsvc)

	mux.HandleFunc("GET /{$}", homeHandler.Home)

	mux.HandleFunc("GET /blog/{$}", blogHandler.ViewAll)
	mux.HandleFunc("GET /blog/view/{slug}", blogHandler.View)
	mux.HandleFunc("GET /blog/create/", blogHandler.Create)
	mux.HandleFunc("POST /blog/create", blogHandler.CreatePost)
	mux.HandleFunc("PUT /blog/update/{slug}", blogHandler.Update)
	mux.HandleFunc("DELETE /blog/{slug}", blogHandler.Delete)

	return middleware.Logger(cfg.Log, mux)
}
