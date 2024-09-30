package models

import (
	"time"

	"github.com/Triyaambak/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func DatabasePostToStructPost(db database.Post) Post {
	var description *string
	if db.Description.Valid {
		description = &db.Description.String
	}
	return Post{
		ID:          db.ID,
		CreatedAt:   db.CreatedAt,
		UpdatedAt:   db.UpdatedAt,
		Title:       db.Title,
		Description: description,
		PublishedAt: db.PublishedAt,
		Url:         db.Url,
		FeedID:      db.FeedID,
	}
}

func DatabsePostsToStructPosts(dbPosts []database.Post) []Post {
	var posts []Post
	for _, post := range dbPosts {
		posts = append(posts, DatabasePostToStructPost(post))
	}
	return posts
}
