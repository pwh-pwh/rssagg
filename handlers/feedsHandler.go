package handlers

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pwh-pwh/rssagg/config"
	"github.com/pwh-pwh/rssagg/internal/database"
	"github.com/pwh-pwh/rssagg/models"
	"github.com/pwh-pwh/rssagg/resp"
	"io/ioutil"
	"net/http"
	"time"
)

func GetFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := config.Config.DB.GetFeeds(context.Background())
	if err != nil {
		resp.RespondWithError(w, 500, err.Error())
		return
	}
	resp.RespondWithJSON(w, 200, models.DbFeedsToFeeds(feeds))
}

func CreateFeedsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type FeedsReqBody struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	reqBody := FeedsReqBody{}
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &reqBody)
	if err != nil {
		resp.RespondWithError(w, 500, "json unmarshal err")
		return
	}
	now := time.Now()
	feed, err := config.Config.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      reqBody.Name,
		Url:       reqBody.Url,
		UserID:    user.ID,
	})
	if err != nil {
		resp.RespondWithError(w, 500, err.Error())
		return
	}
	resp.RespondWithJSON(w, 200, models.DbFeedToFeed(feed))
}
