package logger

import "log/slog"

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

type Config struct {
	Level  Level
	Format Format // text of json
}

type Format string

const (
	FormatText Format = "text"
	FormatJSON Format = "json"
)

func (c Config) slogLevel() slog.Level {
	switch c.Level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func DefaultConfig() Config {
	return Config{
		Level:  LevelInfo,
		Format: FormatText,
	}
}
