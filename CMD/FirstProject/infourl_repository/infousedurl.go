package infourl_repository

import (
	"example.com/gin-project/CMD/FirstProject/db_repository"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

func GetUserData(c *gin.Context, shortUrl string) error {
	// Извлекаем информацию из объекта http.Request
	userAgent := c.Request.UserAgent()
	device := "Device"
	ua := user_agent.New(userAgent)
	os := ua.OS()

	if os == "" {
		os = "Unknown"
	}

	ipAddress := c.ClientIP()

	db, err := db_repository.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	err = db_repository.InsertUserData(db, shortUrl, userAgent, device, os, ipAddress)
	if err != nil {
		return err
	}

	return nil
}
