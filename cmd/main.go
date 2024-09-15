package main

import (
	"fmt"
	"log"
	"net/http"
	"personal/internal/config"
	"personal/internal/handlers"
	"personal/internal/middleware"
	"personal/internal/store/db"
	"personal/internal/store/db_utils"
)

func main() {
	config := config.MustLoadConfig()
	conn := db_utils.MustOpen(config)
	defer conn.Close()
	queries := db.New(conn)

	mux := http.NewServeMux()

	mux.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", handlers.NewIndexHandler(queries).ServeHTTP)

	r := middleware.TextHTMLMiddleware(mux)
	r = middleware.CSPMiddleware(r)
	r = middleware.Logger(r)

	log.Printf("site running on port %s...\n", config.ServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), r)
	log.Panic(err)
}
