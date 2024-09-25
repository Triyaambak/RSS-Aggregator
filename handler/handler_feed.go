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
	"github.com/google/uuid"
)

func HandlerCreateUserFeed(apiCfg *config.ApiConfg, w http.ResponseWriter, r *http.Request) {
	user, err := middleware.AuthMiddleware(apiCfg, r)
	if err != nil {
		RespondWithError(w, 403, fmt.Sprintln("Error while getting user in HandlerCreateUserFeed func", err))
		return
	}
	params := struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{}

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while decoding params in HandlerCreateUserFeed func", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UsersID:   user.ID,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while creating user feed in HandlerCreateUserFeed func", err))
	}

	RespondWithJSON(w, 201, models.DatabaseFeedToStructFeed(feed))
}

func HandlerGetFeed(apiCfg *config.ApiConfg, w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeed(r.Context())
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while getting feeds in HandlerGetFeed func", err))
	}

	RespondWithJSON(w, 201, models.DatabaseFeedsToStructFeeds(feeds))
}
