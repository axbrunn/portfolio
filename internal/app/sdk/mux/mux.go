package mux

import (
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/ui"
)

type Config struct {
	Log *slog.Logger
}

type RouteAdder interface {
	Add(mux *http.ServeMux, cfg Config)
}

func WebAPI(cfg Config, routeAdder RouteAdder) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	routeAdder.Add(mux, cfg)

	return mux
}
