package handlers

import (
	"net/http"
	"personal/internal/store/db"
	"personal/internal/templates"
)

type ContactHandler struct {
	queries *db.Queries
}

func NewContactHandler(queries *db.Queries) *ContactHandler {
	return &ContactHandler{queries: queries}
}

func (handler *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := templates.Layout(nil, "Index").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
