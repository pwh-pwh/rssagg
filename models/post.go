package models

import (
	"github.com/google/uuid"
	"github.com/pwh-pwh/rssagg/internal/database"
	"time"
)

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func DbPost2Post(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title.String,
		Url:         post.Url,
		Description: post.Description.String,
		PublishedAt: post.PublishedAt,
		FeedID:      post.FeedID,
	}
}

func DbPosts2Posts(posts []database.Post) []Post {
	var result []Post
	for _, post := range posts {
		result = append(result, DbPost2Post(post))
	}
	return result
}
