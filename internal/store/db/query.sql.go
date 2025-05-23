// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
)

const getInfo = `-- name: GetInfo :one
SELECT phonenumber, email, user_location, linkedin, github
FROM userinfo
limit 1
`

func (q *Queries) GetInfo(ctx context.Context) (Userinfo, error) {
	row := q.db.QueryRowContext(ctx, getInfo)
	var i Userinfo
	err := row.Scan(
		&i.Phonenumber,
		&i.Email,
		&i.UserLocation,
		&i.Linkedin,
		&i.Github,
	)
	return i, err
}

const getProjectsWSkills = `-- name: GetProjectsWSkills :many
SELECT 
  projects.id, projects.name, projects.description,
  skills.id, skills.name
FROM projects
INNER JOIN projects_skills on projects.id = projects_skills.p_id
INNER JOIN skills on skills.id = projects_skills.s_id
ORDER BY projects.id
`

type GetProjectsWSkillsRow struct {
	Project Project
	Skill   Skill
}

func (q *Queries) GetProjectsWSkills(ctx context.Context) ([]GetProjectsWSkillsRow, error) {
	rows, err := q.db.QueryContext(ctx, getProjectsWSkills)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProjectsWSkillsRow
	for rows.Next() {
		var i GetProjectsWSkillsRow
		if err := rows.Scan(
			&i.Project.ID,
			&i.Project.Name,
			&i.Project.Description,
			&i.Skill.ID,
			&i.Skill.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLanguages = `-- name: ListLanguages :many
SELECT name
FROM languages
`

func (q *Queries) ListLanguages(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listLanguages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProjects = `-- name: ListProjects :many
SELECT id, name, description
FROM projects
`

func (q *Queries) ListProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, listProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSkills = `-- name: ListSkills :many
SELECT id, name
FROM skills
`

func (q *Queries) ListSkills(ctx context.Context) ([]Skill, error) {
	rows, err := q.db.QueryContext(ctx, listSkills)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Skill
	for rows.Next() {
		var i Skill
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSkillsByProjectID = `-- name: ListSkillsByProjectID :many
SELECT skills.id, skills.name
FROM skills
INNER JOIN projects_skills ON skills.id = projects_skills.s_id
INNER JOIN projects ON projects_skills.p_id = projects.id
WHERE projects.id = $1
`

func (q *Queries) ListSkillsByProjectID(ctx context.Context, id int64) ([]Skill, error) {
	rows, err := q.db.QueryContext(ctx, listSkillsByProjectID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Skill
	for rows.Next() {
		var i Skill
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
