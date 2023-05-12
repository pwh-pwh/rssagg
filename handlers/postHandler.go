package handlers

import (
	"context"
	"github.com/pwh-pwh/rssagg/config"
	"github.com/pwh-pwh/rssagg/internal/database"
	"github.com/pwh-pwh/rssagg/models"
	"github.com/pwh-pwh/rssagg/resp"
	"net/http"
	"strconv"
)

func GetPostsLimit(w http.ResponseWriter, r *http.Request, user database.User) {
	limitS := r.URL.Query().Get("limit")
	if limitS == "" {
		limitS = "10"
	}
	limit, err := strconv.Atoi(limitS)
	if err != nil {
		limit = 10
	}
	posts, err := config.Config.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		resp.RespondWithJSON(w, 500, err.Error())
		return
	}
	resp.RespondWithJSON(w, 200, models.DbPosts2Posts(posts))
}
