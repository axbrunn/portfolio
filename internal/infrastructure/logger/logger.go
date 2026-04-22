package logger

import (
	"log/slog"
	"os"
)

func New(cfg Config) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: cfg.slogLevel(),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := a.Value.Time()
				return slog.String(
					slog.TimeKey,
					t.Format("2006/01/02 15:04:05.000"),
				)
			}
			return a
		},
	}

	var handler slog.Handler
	switch cfg.Format {
	case FormatJSON:
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
