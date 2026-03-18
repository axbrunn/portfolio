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

	home := handlers.NewHome(cfg.Log)
	blog := handlers.NewBlog(cfg.Log, blogsvc)

	mux.HandleFunc("GET /{$}", home.Home)
	mux.HandleFunc("GET /blogs", blog.GetAll)
	mux.HandleFunc("GET /blog/{slug}", blog.Get)
	mux.HandleFunc("POST /blog", blog.Insert)
	mux.HandleFunc("PUT /blog/{slug}", blog.Update)
	mux.HandleFunc("DELETE /blog/{slug}", blog.Delete)

	return middleware.Logger(cfg.Log, mux)
}
