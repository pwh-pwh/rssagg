package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/pwh-pwh/rssagg/internal/database"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

var ApiConfig apiConfig

func main() {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("can not load env:port")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("db url is null")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can not open db_url")
	}
	queries := database.New(db)
	ApiConfig.DB = queries
	//Create a chi.NewRouter
	router := chi.NewRouter()
	//Use router.Use to add the built-in cors.Handler middleware.
	router.Use(cors.AllowAll().Handler)
	Router(router)
	err = http.ListenAndServe(":"+portStr, router)
	if err != nil {
		return
	}
}
