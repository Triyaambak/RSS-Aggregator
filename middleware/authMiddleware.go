package middleware

import (
	"net/http"

	auth "github.com/Triyaambak/RSS-Aggregator/auth"
	config "github.com/Triyaambak/RSS-Aggregator/config"
	"github.com/Triyaambak/RSS-Aggregator/internal/database"
)

func AuthMiddleware(apiCfg *config.ApiConfg, r *http.Request) (database.User, error) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		return database.User{}, err
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}
