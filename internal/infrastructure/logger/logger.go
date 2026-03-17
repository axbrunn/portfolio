package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
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

	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(handler)
}
