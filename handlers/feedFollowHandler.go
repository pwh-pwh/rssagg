package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pwh-pwh/rssagg/config"
	"github.com/pwh-pwh/rssagg/internal/database"
	"github.com/pwh-pwh/rssagg/models"
	"github.com/pwh-pwh/rssagg/resp"
	"io/ioutil"
	"net/http"
	"time"
)

func GetAllFF4UserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := config.Config.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		resp.RespondWithError(w, 400, err.Error())
		return
	}
	resp.RespondWithJSON(w, 200, models.DbFFs2FFs(feedFollows))
}

func DelFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId := chi.URLParam(r, "feedFollowID")
	fFId, err := uuid.Parse(feedFollowId)
	if err != nil {
		resp.RespondWithError(w, 400, err.Error())
		return
	}
	err = config.Config.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		ID:     fFId,
		UserID: user.ID,
	})
	if err != nil {
		resp.RespondWithError(w, 500, err.Error())
		return
	}
	resp.RespondWithJSON(w, 200, "del feed_follow ok")
}

func CreateFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type ReqBody struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	var reqBody ReqBody
	all, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(all, &reqBody)
	if err != nil {
		resp.RespondWithError(w, 400, err.Error())
		return
	}
	var now time.Time = time.Now()
	feedFollow, err := config.Config.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    reqBody.FeedId,
	})
	if err != nil {
		resp.RespondWithError(w, 500, err.Error())
		return
	}
	resp.RespondWithJSON(w, 200, models.DbFF2FeedFollow(feedFollow))
}
