package main

import (
	"log"

	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/config"
	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/controller"
	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/db"
	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/logic"
)

func main() {
	cfg := config.Read()
	db, err := db.New(cfg.DBCon, cfg.SiteURL)
	if err != nil {
		log.Fatalf("Ошибка при создании объекта бд - %s", err.Error())
	}
	logic := logic.New(cfg, &db)
	controller.New(cfg, logic).Start()
}
