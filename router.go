package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pwh-pwh/rssagg/internal/database"
	"io/ioutil"
	"net/http"
	"time"
)

func Router(router chi.Router) {
	//Create sub-router for the /v1 namespace and mount it to the main router.
	router.Route("/v1", func(r chi.Router) {
		r.Get("/readiness", func(writer http.ResponseWriter, request *http.Request) {
			respondWithJSON(writer, 200, struct {
				Status string `json:"status"`
			}{
				"ok",
			})
		})
		r.Get("/err", func(writer http.ResponseWriter, request *http.Request) {
			respondWithError(writer, 500, "Internal Server Error")
		})
		r.Post("/users", func(writer http.ResponseWriter, request *http.Request) {
			type Body struct {
				Name string `json:"name"`
			}
			var body Body
			all, _ := ioutil.ReadAll(request.Body)
			_ = json.Unmarshal(all, &body)
			newUUID, _ := uuid.NewUUID()
			var now time.Time
			user, err := ApiConfig.DB.CreateUser(context.Background(), database.CreateUserParams{
				ID:        newUUID,
				CreatedAt: now,
				UpdatedAt: now,
				Name:      body.Name,
			})
			if err != nil {
				respondWithError(writer, 500, err.Error())
			}
			respondWithJSON(writer, 200, user)
		})
	})
}
