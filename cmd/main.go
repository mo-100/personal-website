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
	cfg := config.MustLoadConfig()
	conn := db_utils.GetDB(cfg)
	defer conn.Close()
	queries := db.New(conn)

	mux := http.NewServeMux()

	mux.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", handlers.NewIndexHandler(queries).ServeHTTP)
	mux.HandleFunc("/blog", handlers.NewBlogHandler(queries).ServeHTTP)
	mux.HandleFunc("/cv", handlers.NewCVHandler(queries).ServeHTTP)
	mux.HandleFunc("/contactme", handlers.NewContactHandler(queries).ServeHTTP)
	mux.HandleFunc("/portfolio", handlers.NewPortfolioHandler(queries).ServeHTTP)

	r := middleware.TextHTMLMiddleware(mux)
	r = middleware.CSPMiddleware(r)
	r = middleware.Logger(r)

	log.Printf("site running on port %s...\n", cfg.ServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServerPort), r)
	log.Panic(err)
}
