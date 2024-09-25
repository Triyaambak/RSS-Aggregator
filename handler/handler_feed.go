package handler

import (
	"net/http"

	config "github.com/Triyaambak/RSS-Aggregator/config"
	"github.com/Triyaambak/RSS-Aggregator/internal/database"
)

func HadlerCreateUserFeed(apiCfg *config.ApiConfg, w http.ResponseWriter, r *http.Request, user database.User) {

}
