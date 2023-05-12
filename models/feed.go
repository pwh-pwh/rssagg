package models

import (
	"github.com/google/uuid"
	"github.com/pwh-pwh/rssagg/internal/database"
	"time"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func DbFeedToFeed(feedDb database.Feed) Feed {
	return Feed{
		ID:        feedDb.ID,
		CreatedAt: feedDb.CreatedAt,
		UpdatedAt: feedDb.UpdatedAt,
		Name:      feedDb.Name,
		Url:       feedDb.Url,
		UserID:    feedDb.UserID,
	}
}

func DbFeedsToFeeds(fds []database.Feed) []Feed {
	result := make([]Feed, 0, len(fds))
	for _, f := range fds {
		result = append(result, DbFeedToFeed(f))
	}
	return result
}
