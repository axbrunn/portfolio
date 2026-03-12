package build

import (
	"net/http"

	"github.com/axbrunn/portfolio/internal/app/domain/blogapp"
	"github.com/axbrunn/portfolio/internal/app/domain/homeapp"
	"github.com/axbrunn/portfolio/internal/app/sdk/mux"
)

type routes struct{}

func Routes() *routes {
	return &routes{}
}

func (r *routes) Add(m *http.ServeMux, cfg mux.Config) {
	homeapp.Routes(m, homeapp.Config{
		Log: cfg.Log,
	})

	blogapp.Routes(m, blogapp.Config{
		Log: cfg.Log,
	})
}
