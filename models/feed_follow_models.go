package models

import (
	"time"

	"github.com/Triyaambak/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

func DatabaseFeedFollowToStructFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		FeedID:    dbFeedFollow.FeedID,
		UserID:    dbFeedFollow.UserID,
	}
}
