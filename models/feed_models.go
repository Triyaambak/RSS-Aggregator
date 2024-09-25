package models

import (
	"time"

	"github.com/Triyaambak/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UsersID   uuid.UUID `json:"users_id"`
}

func DatabaseFeedToStructFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UsersID:   dbFeed.UsersID,
	}
}

func DatabaseFeedsToStructFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, 0)
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToStructFeed(dbFeed))
	}
	return feeds
}
