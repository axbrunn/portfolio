package bloghandler

import (
	"log/slog"
	"net/http"
)

type Config struct {
	Log     *slog.Logger
	Service BlogService
}

func Routes(mux *http.ServeMux, cfg Config) {
	h := &handlers{
		log:     cfg.Log,
		service: cfg.Service,
	}

	mux.HandleFunc("GET /blogs", h.getAll)
	mux.HandleFunc("GET /blog/{slug}", h.get)
	mux.HandleFunc("POST /blog", h.insert)
	mux.HandleFunc("PUT /blog/{slug}", h.update)
	mux.HandleFunc("DELETE /blog/{slug}", h.delete)
}
