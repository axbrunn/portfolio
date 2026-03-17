package build

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/internal/handler/bloghandler"
	"github.com/axbrunn/portfolio/internal/handler/homehandler"
	"github.com/axbrunn/portfolio/internal/infrastructure/mux"
	"github.com/axbrunn/portfolio/internal/repository/blogrepo"
	"github.com/axbrunn/portfolio/internal/service/blogservice"
)

type Config struct {
	Log *slog.Logger
	DB  *sql.DB
}

type routes struct {
	cfg Config
}

func Routes(cfg Config) *routes {
	return &routes{cfg: cfg}
}

func (r *routes) Add(m *http.ServeMux, _ mux.Config) {
	blogRepo := blogrepo.New(r.cfg.DB)
	blogSvc := blogservice.New(blogRepo)

	homehandler.Routes(m, homehandler.Config{
		Log: r.cfg.Log,
	})

	bloghandler.Routes(m, bloghandler.Config{
		Log:     r.cfg.Log,
		Service: blogSvc,
	})
}
