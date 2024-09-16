package handlers

import (
	"fmt"
	"net/http"
	"personal/internal/store/db"
	"personal/internal/templates"
	"personal/internal/utils"
)

type IndexHandler struct {
	queries *db.Queries
}

func NewIndexHandler(queries *db.Queries) *IndexHandler {
	return &IndexHandler{queries: queries}
}

func (handler *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projects_skills, err := handler.queries.GetProjectsWSkills(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting projects: %v", err), http.StatusInternalServerError)
		return
	}
	p := utils.Map(projects_skills, func(x db.GetProjectsWSkillsRow) db.Project { return x.Project })
	s := utils.Map(projects_skills, func(x db.GetProjectsWSkillsRow) db.Skill { return x.Skill })

	p2, s2 := utils.GetOneToMany(p, s, func(x db.Project) string { return string(x.ID) })

	c := templates.Projects(p2, s2)
	err = templates.Layout(c, "Index").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
