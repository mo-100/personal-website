package handlers

import (
	"net/http"
	"personal/internal/store/db"
	"personal/internal/templates"
)

type PortfolioHandler struct {
	queries *db.Queries
}

func NewPortfolioHandler(queries *db.Queries) *PortfolioHandler {
	return &PortfolioHandler{queries: queries}
}

func (handler *PortfolioHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := templates.Layout(templates.Empty(), "Index").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
