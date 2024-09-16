package handlers

import (
	"net/http"
	"personal/internal/store/db"
	"personal/internal/templates"
)

type CVHandler struct {
	queries *db.Queries
}

func NewCVHandler(queries *db.Queries) *CVHandler {
	return &CVHandler{queries: queries}
}

func (handler *CVHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := templates.Layout(nil, "Index").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
