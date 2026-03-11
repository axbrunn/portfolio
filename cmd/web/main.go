package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/axbrunn/portfolio/cmd/web/build"
	"github.com/axbrunn/portfolio/internal/app/sdk/mux"
	"github.com/axbrunn/portfolio/internal/foundation/config"
	"github.com/axbrunn/portfolio/internal/foundation/logger"
	"github.com/axbrunn/portfolio/internal/foundation/web"
)

func main() {
	cfg := config.LoadConfig()
	logger := logger.New()
	slog.SetDefault(logger)

	webAPI := mux.WebAPI(mux.Config{
		Log: logger,
	}, build.Routes())

	srv := web.NewServer(web.Config{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      webAPI,
		Logger:       logger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	err := srv.Start()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
