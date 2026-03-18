package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/axbrunn/portfolio/internal/application/routes"
	"github.com/axbrunn/portfolio/internal/infrastructure/config"
	"github.com/axbrunn/portfolio/internal/infrastructure/database"
	"github.com/axbrunn/portfolio/internal/infrastructure/logger"
	"github.com/axbrunn/portfolio/internal/infrastructure/web"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.New()
	slog.SetDefault(log)

	db, err := database.Open(cfg.DB)
	if err != nil {
		slog.Error("database connection failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	webAPI := routes.New(routes.Config{
		Log: log,
		DB:  db,
	})

	srv := web.NewServer(web.Config{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      webAPI,
		Logger:       log,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	if err := srv.Start(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
