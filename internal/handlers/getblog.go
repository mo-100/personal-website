package handlers

import (
	"net/http"
	"personal/internal/store/db"
	"personal/internal/templates"
)

type BlogHandler struct {
	queries *db.Queries
}

func NewBlogHandler(queries *db.Queries) *BlogHandler {
	return &BlogHandler{queries: queries}
}

func (handler *BlogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := templates.Layout(nil, "Index").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
