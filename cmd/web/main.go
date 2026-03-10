package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/axbrunn/portfolio/internal/app"
	"github.com/axbrunn/portfolio/internal/config"
	"github.com/axbrunn/portfolio/internal/logger"
	"github.com/axbrunn/portfolio/internal/web"
)

func main() {
	cfg := config.LoadConfig()
	logger := logger.New()
	slog.SetDefault(logger)

	// logger.Info("connecting to database")

	// db, err := app.OpenDB(*cfg)
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	os.Exit(1)
	// }
	//
	// defer func() {
	// 	logger.Info("closing database connection")
	// 	db.Close()
	// }()
	//
	// logger.Info("database connection pool established")

	app := &app.Application{
		Logger: logger,
		Config: cfg,
	}

	srv := web.NewServer(web.Config{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      web.Routes(app),
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
