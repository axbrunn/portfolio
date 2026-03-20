package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/axbrunn/portfolio/internal/application/handlers"
	"github.com/axbrunn/portfolio/internal/application/routes"
	"github.com/axbrunn/portfolio/internal/infrastructure/config"
	"github.com/axbrunn/portfolio/internal/infrastructure/database"
	"github.com/axbrunn/portfolio/internal/infrastructure/logger"
	"github.com/axbrunn/portfolio/internal/infrastructure/web"

	blogbus "github.com/axbrunn/portfolio/internal/business/blog"
	blogrepo "github.com/axbrunn/portfolio/internal/store/blog"
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

	blogrepo := blogrepo.New(db)
	blogsvc := blogbus.New(blogrepo)

	routes := routes.New(routes.Router{
		Logger:      log,
		HomeHandler: handlers.NewHome(log),
		BlogHandler: handlers.NewBlog(log, blogsvc),
	})

	srv := web.NewServer(web.Config{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      routes,
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
