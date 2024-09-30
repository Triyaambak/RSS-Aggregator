package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	config "github.com/Triyaambak/RSS-Aggregator/config"
	"github.com/Triyaambak/RSS-Aggregator/internal/database"
	middleware "github.com/Triyaambak/RSS-Aggregator/middleware"
	models "github.com/Triyaambak/RSS-Aggregator/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func HandlerCreateFeedFollow(apiCfg *config.ApiConfg, w http.ResponseWriter, r *http.Request) {
	user, err := middleware.AuthMiddleware(apiCfg, r)
	if err != nil {
		RespondWithError(w, 403, fmt.Sprintln("Error while getting user in HandlerCreateFeedFollow func", err))
		return
	}
	params := struct {
		FeedID uuid.UUID `json:"feed_id"`
	}{}

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while decoding params in HandlerCreateFeedFollow func", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while creating feed follow in HandlerCreateFeedFollow func", err))
	}

	RespondWithJSON(w, 201, models.DatabaseFeedFollowToStructFeedFollow(feed))
}

func HandlerGetFeedFollows(apiCfg *config.ApiConfg, w http.ResponseWriter, r *http.Request) {
	user, err := middleware.AuthMiddleware(apiCfg, r)
	if err != nil {
		RespondWithError(w, 403, fmt.Sprintln("Error while getting user in HandlerGetFeedFollows func", err))
		return
	}

	feed, err := apiCfg.DB.GetFeedFollowsByUserId(r.Context(), user.ID)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while creating feed follow in HandlerGetFeedFollows func", err))
	}

	RespondWithJSON(w, 201, models.DatabseFeedFollowsToStructFeedFollows(feed))
}

func HandlerDeleteFeedFollow(apiCfg *config.ApiConfg, w http.ResponseWriter, r *http.Request) {
	user, err := middleware.AuthMiddleware(apiCfg, r)
	if err != nil {
		RespondWithError(w, 403, fmt.Sprintln("Error while getting user in HandlerDeleteFeedFollow func", err))
		return
	}

	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while parsing feed follow id in HandlerDeleteFeedFollow func", err))
		return
	}

	err = apiCfg.DB.DeletFeedFollow(r.Context(), database.DeletFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while delete feed follow in HandlerDeleteFeedFollow func", err))
	}

	RespondWithJSON(w, 200, "Feed follow deleted successfully")
}
