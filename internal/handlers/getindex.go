package handlers

import (
	"net/http"
	"personal/internal/store/db"
	"personal/internal/templates"
)

type IndexHandler struct {
	queries *db.Queries
}

func NewIndexHandler(queries *db.Queries) *IndexHandler {
	return &IndexHandler{queries: queries}
}

func (handler *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := templates.Layout(templates.Home(), "Home").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
