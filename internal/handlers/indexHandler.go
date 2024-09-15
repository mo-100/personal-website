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

func (ih *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projects, err := ih.queries.ListProjects(r.Context())
	if err != nil {
		http.Error(w, "Error getting projects", http.StatusInternalServerError)
		return
	}
	c := templates.Projects(projects)
	err = templates.Layout(c, "Index").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
