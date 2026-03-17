package web

import (
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
}

type Config struct {
	Addr         string
	Handler      http.Handler
	Logger       *slog.Logger
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServer(cfg Config) *Server {
	s := &Server{
		logger: cfg.Logger,
		httpServer: &http.Server{
			Addr:         cfg.Addr,
			Handler:      cfg.Handler,
			IdleTimeout:  cfg.IdleTimeout,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			ErrorLog:     slog.NewLogLogger(cfg.Logger.Handler(), slog.LevelError),
		},
	}

	return s
}

func (s *Server) Start() error {
	s.logger.Info("server started", "addr", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
