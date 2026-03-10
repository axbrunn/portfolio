package app

import (
	"log/slog"

	"github.com/axbrunn/portfolio/internal/config"
)

type Application struct {
	Logger *slog.Logger
	Config *config.Config
}
