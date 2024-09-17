package handlers

import (
	"fmt"
	"net/http"
	"personal/internal/store/db"
	"personal/internal/templates"
	"personal/internal/utils"
)

type CVHandler struct {
	queries *db.Queries
}

func NewCVHandler(queries *db.Queries) *CVHandler {
	return &CVHandler{queries: queries}
}

func (handler *CVHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projectsSkills, err := handler.queries.GetProjectsWSkills(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting projects: %v", err), http.StatusInternalServerError)
		return
	}
	p, s := utils.GetOneToMany(
		utils.Map(projectsSkills, func(x db.GetProjectsWSkillsRow) db.Project { return x.Project }),
		utils.Map(projectsSkills, func(x db.GetProjectsWSkillsRow) db.Skill { return x.Skill }),
		func(x db.Project) int64 { return x.ID },
	)

	c := templates.CV(p, s)
	err = templates.Layout(c, "CV").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
