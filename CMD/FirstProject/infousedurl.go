package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

type UserData struct {
	UserAgent string
	Device    string
	OS        string
	IP        string
}

func getUserData(c *gin.Context, shortUrl string) error {
	// Извлекаем информацию из объекта http.Request
	userAgent := c.Request.UserAgent()
	device := "Device"
	ua := user_agent.New(userAgent)
	os := ua.OS()

	if os == "" {
		os = "Unknown"
	}

	ipAddress := c.ClientIP()

	userData := &UserData{
		UserAgent: userAgent,
		Device:    device,
		OS:        os,
		IP:        ipAddress,
	}

	db, err := connectToPostgreSQL()
	if err != nil {
		return err
	}

	err = insertUserData(db, shortUrl, userData)
	if err != nil {
		return err
	}

	return nil
}
