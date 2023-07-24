package main

import (
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

	db, err := ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	err = InsertUserData(db, shortUrl, userAgent, device, os, ipAddress)
	if err != nil {
		return err
	}

	return nil
}
