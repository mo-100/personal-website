package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

var db *sql.DB

type ProfileDbRow struct {
	Name      string
	Education string
	Summary   string
}

func initDatabase(dbPath string) error {

	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS profile (
			name TEXT NOT NULL, 
			education TEXT NOT NULL, 
			summary TEXT NOT NULL
		)`,
	)
	if err != nil {
		return err
	}
	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	initDatabase("sqlite.db")
	mux := http.NewServeMux()

	mux.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", indexHandler)

	log.Println("site running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
