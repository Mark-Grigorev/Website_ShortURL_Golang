package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	statususedurl "example.com/gin-project/CMD/FirstProject/infourl_repository"
	shorturl "example.com/gin-project/CMD/FirstProject/shorturl_repository"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Добавляем middleware для поддержки CORS
	// Без этого не работало, еще не разобрался что и почему :(
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.POST("/urlshort", func(c *gin.Context) {
		//Метод создания короткой ссылки
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("Cannot read body"))
		}

		shortURL, err := shorturl.GenerateshortURL(string(body)) // Генерируем короткий URL из длинного URL
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.String(http.StatusOK, shortURL)
	})

	r.GET("/:shortID", func(c *gin.Context) {
		shortID := c.Param("shortID")
		longURL, err := shorturl.FindshortUrl(shortID)
		if err != nil {
			c.String(http.StatusNotFound, "URL not found", err)
			return
		} else {
			//Вызов функции, которая считывает данные пользователя
			err := statususedurl.GetUserData(c, shortID)
			if err != nil {
				fmt.Println(err)
			}
			//Редирект пользователя по длинной ссылке
			c.Redirect(http.StatusMovedPermanently, longURL)
		}
	})

	r.POST("/statusurl", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("Cannot read body"))
		}
		statusUrl, err := shorturl.StatusShortUrl(string(body))
		if err != nil {
			c.AbortWithError(http.StatusNotFound, errors.New("Not found short url"))
		}
		c.String(http.StatusOK, statusUrl)
	})

	return r
}

func main() {
	localhost := ":8080"
	r := setupRouter()
	r.Run(localhost)
}
