package handlers

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pwh-pwh/rssagg/config"
	"github.com/pwh-pwh/rssagg/internal/auth"
	"github.com/pwh-pwh/rssagg/internal/database"
	"github.com/pwh-pwh/rssagg/models"
	"github.com/pwh-pwh/rssagg/resp"
	"io/ioutil"
	"net/http"
	"time"
)

func CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	type Body struct {
		Name string `json:"name"`
	}
	var body Body
	all, _ := ioutil.ReadAll(request.Body)
	_ = json.Unmarshal(all, &body)
	newUUID, _ := uuid.NewUUID()
	var now time.Time
	user, err := config.Config.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        newUUID,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      body.Name,
	})
	if err != nil {
		resp.RespondWithError(writer, 500, err.Error())
	}
	resp.RespondWithJSON(writer, 200, models.DatabaseUserToUser(user))
}

func GetUserHandler(w http.ResponseWriter, req *http.Request) {
	apiKey, err := auth.GetAPIKey(req.Header)
	if err != nil {
		resp.RespondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
		return
	}
	user, err := config.Config.DB.GetUserByApiKey(context.Background(), apiKey)
	if err != nil {
		resp.RespondWithError(w, 400, "api_key err")
		return
	}
	resp.RespondWithJSON(w, 200, models.DatabaseUserToUser(user))
}