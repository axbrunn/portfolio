package homeapp

import (
	"log/slog"
	"net/http"
)

type Config struct {
	Log *slog.Logger
}

func Routes(mux *http.ServeMux, cfg Config) {
	h := &handlers{log: cfg.Log}

	mux.HandleFunc("GET /home", h.home)
}
