package config

import (
	"log"
	"os"

	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/model"
)

func Read() *model.Config {
	var cfg model.Config

	cfg.DBCon = os.Getenv("DB_CONNECTION_STRING")
	if cfg.DBCon == "" {
		log.Fatalf("Не указан DB_CONN")
	}

	cfg.Host = os.Getenv("HOST")
	if cfg.Host == "" {
		cfg.Host = "8080"
	}

	cfg.SiteURL = os.Getenv("SITE_URL")
	if cfg.SiteURL == "" {
		log.Fatalf("Не указан SITE_URL")
	}

	return &cfg
}
