package controller

import (
	"log"
	"net/http"

	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/model"
	"github.com/gin-gonic/gin"
)

func (a *App) setURLShorterRoutes(group *gin.RouterGroup) {
	group.POST("/create", a.CreateURLShort)
	group.POST("/status", a.StatusURL)
	group.GET("/:shortID", a.RedirectURL)
}

func (a *App) CreateURLShort(ctx *gin.Context) {
	var err error
	var shortURL string

	userID, ok := ctx.Get("user_id")
	if !ok {
		a.Error(ctx, http.StatusUnauthorized, "hasn't_user_id_in_ctx")
		return
	}
	userIDInt, _ := userID.(int64)

	var URLReq model.URLReq
	if err = ctx.ShouldBindJSON(&URLReq); err != nil {
		a.Error(ctx, http.StatusBadRequest, "некорректные данные")
		return
	}

	if URLReq.URL == "" {
		a.Error(ctx, http.StatusBadRequest, "не указана ссылка")
		return
	}

	// Генерируем короткий URL из длинного URL
	shortURL, err = a.logic.GetOrCreateShortURL(ctx.Request.Context(), URLReq.URL, userIDInt)
	if err != nil {
		a.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, shortURL)
}

func (a *App) StatusURL(ctx *gin.Context) {
	var err error

	var URLReq model.URLReq
	if err = ctx.ShouldBindJSON(&URLReq); err != nil {
		a.Error(ctx, http.StatusBadRequest, "некорректные данные")
	}
	if URLReq.ID == "" {
		a.Error(ctx, http.StatusBadRequest, "не указан id")
		return
	}
	info, err := a.logic.StatusShortUrl(ctx.Request.Context(), URLReq.ID)
	if err != nil {
		a.Error(ctx, http.StatusNotFound, "not found short url")
		return
	}
	ctx.JSON(http.StatusOK, info)
}

func (a *App) RedirectURL(ctx *gin.Context) {
	shortID := ctx.Param("shortID")
	longURL, err := a.logic.FindLongUrl(ctx.Request.Context(), shortID)
	if err != nil {
		a.Error(ctx, http.StatusNotFound, "URL not found")
		return
	} else {
		// Реддиректим пользователя и параллельно собираем данные.
		go func() {
			err := a.logic.GetUserData(ctx, shortID)
			if err != nil {
				log.Fatalln(err)
			}
		}()
		ctx.Redirect(http.StatusMovedPermanently, longURL)
	}
}
