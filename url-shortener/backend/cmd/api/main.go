package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/z66x/url-shortener/internal/db"
	"github.com/z66x/url-shortener/internal/handlers"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("error loading .env file")
	}

	database := db.New()
	defer database.Close()

	urlRepo := db.NewURLRepository(database)
	urlHandler := handlers.NewURLHandler(urlRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/shorten", urlHandler.Shorten)
	r.Get("/{code}", urlHandler.Redirect)

	port := os.Getenv("APP_PORT")
	log.Printf("server running on port %s", port)
	http.ListenAndServe(":"+port, r)
}