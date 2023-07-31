package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorizationData struct {
	Login string `json:"Login"`
	Name  string `json:"Password"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Добавляем middleware для поддержки CORS

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

		shortURL, err := GenerateshortURL(string(body)) // Генерируем короткий URL из длинного URL
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.String(http.StatusOK, shortURL)
	})

	r.POST("/statusurl", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("Cannot read body"))
			return
		}
		shortID := string(body)
		info, err := StatusShortUrl(shortID)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, errors.New("Not found short url"))
			return
		}
		c.JSON(http.StatusOK, info)
	})

	r.POST("/registration", func(ctx *gin.Context) {
		var messageRegistr string
		messageRegistr = "Вы успешно зарегестрировались"
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("Cannot read body"))
			return
		}

		err = registration(string(body))
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("Error registration"))
			return
		}

		ctx.JSON(http.StatusOK, messageRegistr)
	})

	r.POST("/authorization", func(ctx *gin.Context) {
		var messageAuth string
		messageAuth = "Вы успешно авторизировались"

		var authData AuthorizationData
		if err := ctx.ShouldBindJSON(&authData); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("некорректные данные"))
			return
		}

		err := authorization(authData.Login, authData.Name)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("ошибка авторизации"))
			return
		}

		ctx.JSON(http.StatusOK, messageAuth)

	})

	r.GET("/:shortID", func(c *gin.Context) {
		shortID := c.Param("shortID")
		longURL, err := FindshortUrl(shortID)
		if err != nil {
			c.String(http.StatusNotFound, "URL not found", err)
			return
		} else {
			//Вызов функции, которая считывает данные пользователя
			err := GetUserData(c, shortID)
			if err != nil {
				fmt.Println(err)
			}
			//Редирект пользователя по длинной ссылке
			c.Redirect(http.StatusMovedPermanently, longURL)
		}
	})

	return r
}

func main() {
	localhost := ":8080"
	r := setupRouter()
	r.Run(localhost)
}
