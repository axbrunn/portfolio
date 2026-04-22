package logger

import (
	"context"
	"log/slog"
)

type Handler struct {
	logger *slog.Logger
}

func NewHandler(logger *slog.Logger) *Handler {
	return &Handler{logger: logger}
}

// Context-aware — handig voor trace/request IDs
func (h *Handler) WithContext(ctx context.Context) *slog.Logger {
	// Straks kun je hier bijv. een requestID uit de context halen
	// requestID := ctx.Value(RequestIDKey)
	return h.logger
}

func (h *Handler) Info(msg string, args ...any) {
	h.logger.Info(msg, args...)
}

func (h *Handler) Warn(msg string, args ...any) {
	h.logger.Warn(msg, args...)
}

func (h *Handler) Error(msg string, args ...any) {
	h.logger.Error(msg, args...)
}

func (h *Handler) Debug(msg string, args ...any) {
	h.logger.Debug(msg, args...)
}

// With voegt vaste velden toe — handig per service/component
func (h *Handler) With(args ...any) *Handler {
	return &Handler{logger: h.logger.With(args...)}
}
