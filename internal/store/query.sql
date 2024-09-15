-- name: ListProjects :many
SELECT *
FROM projects;

-- name: ListSkills :many
SELECT *
FROM skills;

-- name: ListSkillsByProjectID :many
SELECT skills.*
FROM skills
INNER JOIN projects_skills ON skills.id = projects_skills.s_id
INNER JOIN projects ON projects_skills.p_id = projects.id
WHERE projects.id = $1;

-- name: ListLanguages :many
SELECT *
FROM languages;

-- name: GetInfo :one
SELECT *
FROM userinfo
limit 1;