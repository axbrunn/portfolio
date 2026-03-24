package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/axbrunn/portfolio/internal/application/helpers"
	"github.com/axbrunn/portfolio/internal/business/blog"
	"github.com/axbrunn/portfolio/ui/html/pages"
)

type Blog struct {
	logger *slog.Logger
	svc    blog.Service
}

func NewBlog(log *slog.Logger, svc blog.Service) *Blog {
	return &Blog{
		logger: log,
		svc:    svc,
	}
}

func (h *Blog) ViewAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.svc.ViewAll(r.Context())
	if err != nil {
		helpers.ServerError(h.logger, w, r, err)
		return
	}

	for _, post := range posts {
		fmt.Fprintf(w, "%+v\n", post)
	}
}

func (h *Blog) View(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	post, err := h.svc.View(r.Context(), slug)
	if err != nil {
		if errors.Is(err, blog.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			helpers.ServerError(h.logger, w, r, err)
		}
		return
	}

	pages.BlogView(post.Title).Render(r.Context(), w)
}

func (h *Blog) Create(w http.ResponseWriter, r *http.Request) {
	pages.BlogCreate(pages.BlogCreateForm{}).Render(r.Context(), w)
}

func (h *Blog) CreatePost(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 2<<20)

	var form pages.BlogCreateForm
	err := helpers.DecodePostForm(r, &form)

	slug, err := h.svc.CreatePost(r.Context(), form.Title, form.Slug, form.Excerpt, form.Body, form.Published)
	if err != nil {
		// Als het validatiefouten zijn, form opnieuw tonen met de fouten per veld.
		var valErr *blog.ValidationError
		if errors.As(err, &valErr) {
			form.FieldErrors = valErr.Fields
			w.WriteHeader(http.StatusUnprocessableEntity)
			pages.BlogCreate(form).Render(r.Context(), w)
			return
		}
		helpers.ServerError(h.logger, w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/blog/view/%s", slug), http.StatusSeeOther)
}

func (h *Blog) Update(w http.ResponseWriter, r *http.Request) {
}

func (h *Blog) UpdatePut(w http.ResponseWriter, r *http.Request) {
}

func (h *Blog) Delete(w http.ResponseWriter, r *http.Request) {
}
