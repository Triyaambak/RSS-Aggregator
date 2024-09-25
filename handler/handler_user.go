package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	config "github.com/Triyaambak/RSS-Aggregator/config"
	"github.com/Triyaambak/RSS-Aggregator/internal/database"
	middleware "github.com/Triyaambak/RSS-Aggregator/middleware"
	"github.com/google/uuid"
)

func HandlerCreateUser(apiCfg *config.ApiConfg, w http.ResponseWriter, r *http.Request) {
	params := struct {
		Name string `json:"name"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while decoding params in HandlerCreateUser func", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while creating user in HandlerCreateUser func", err))
	}

	RespondWithJSON(w, 201, user)
}

func HandlerGetUser(apiCfg *config.ApiConfg, w http.ResponseWriter, r *http.Request) {
	user, err := middleware.AuthMiddleware(apiCfg, r)
	if err != nil {
		RespondWithError(w, 403, fmt.Sprintln("Error while getting user in HandlerGetUser func", err))
	}
	RespondWithJSON(w, 200, user)
}
