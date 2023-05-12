package models

import (
	"github.com/google/uuid"
	"github.com/pwh-pwh/rssagg/internal/database"
	"time"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"userID"`
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
