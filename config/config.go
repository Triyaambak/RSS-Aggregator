package config

import (
	"database/sql"
	"log"

	"github.com/Triyaambak/RSS-Aggregator/internal/database"
)

type ApiConfg struct {
	DB *database.Queries
}

func ConnectDB(dbUrl string) *ApiConfg {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := ApiConfg{
		DB: database.New(db),
	}

	return &apiCfg
}
