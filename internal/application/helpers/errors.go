package helpers

import (
	"log/slog"
	"net/http"
)

func ServerError(logger *slog.Logger, w http.ResponseWriter, r *http.Request, err error) {
	logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
