package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/pwh-pwh/rssagg/config"
	"github.com/pwh-pwh/rssagg/scraper"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("can not load env:port")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("db url is null")
	}
	err := config.InitConfig(dbURL)
	if err != nil {
		log.Fatal("init db err")
	}
	//Create a chi.NewRouter
	router := chi.NewRouter()
	//Use router.Use to add the built-in cors.Handler middleware.
	router.Use(cors.AllowAll().Handler)
	Router(router)
	const collectionConcurrency = 10
	const collectionInterval = time.Second * 60
	go scraper.StartScraping(collectionConcurrency, collectionInterval)
	err = http.ListenAndServe(":"+portStr, router)
	if err != nil {
		return
	}
}
