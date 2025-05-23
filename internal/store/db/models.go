// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

type Education struct {
	ID   int64
	Name string
}

type Language struct {
	Name string
}

type Project struct {
	ID          int64
	Name        string
	Description string
}

type ProjectsSkill struct {
	PID int64
	SID int64
}

type Skill struct {
	ID   int64
	Name string
}

type Userinfo struct {
	Phonenumber  string
	Email        string
	UserLocation string
	Linkedin     string
	Github       string
}
