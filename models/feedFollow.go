package models

import (
	"github.com/google/uuid"
	"github.com/pwh-pwh/rssagg/internal/database"
	"time"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func DbFF2FeedFollow(ff database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        ff.ID,
		CreatedAt: ff.CreatedAt,
		UpdatedAt: ff.UpdatedAt,
		UserID:    ff.UserID,
		FeedID:    ff.FeedID,
	}
}

func DbFFs2FFs(ffs []database.FeedFollow) []FeedFollow {
	var result []FeedFollow
	for _, f := range ffs {
		result = append(result, DbFF2FeedFollow(f))
	}
	return result
}
