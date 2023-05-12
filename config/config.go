package config

import (
	"database/sql"
	"github.com/pwh-pwh/rssagg/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

var Config ApiConfig

func InitConfig(dbUrl string) error {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return err
	}
	queries := database.New(db)
	Config.DB = queries
	return nil
}
