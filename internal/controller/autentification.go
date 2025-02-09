package controller

import (
	"net/http"

	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/model"
	"github.com/gin-gonic/gin"
)

func (a *App) setAuthRoutes(group *gin.RouterGroup) {
	// TODO - вынести в отдельный сервис
	group.POST("/registration", a.Registration)
	group.POST("/authorization", a.Autorization)
}

func (a *App) Registration(ctx *gin.Context) {
	var err error
	messageRegistr := "Вы успешно зарегестрировались"

	var authData model.AuthorizationData
	if err = ctx.ShouldBindJSON(&authData); err != nil {
		a.Error(ctx, http.StatusBadRequest, "некорректные данные")
		return
	}

	err = a.logic.Authorization(ctx.Request.Context(), authData.Login, authData.Password)
	if err != nil {
		a.Error(ctx, http.StatusBadRequest, "error registration")
		return
	}

	ctx.JSON(http.StatusOK, messageRegistr)
}

func (a *App) Autorization(ctx *gin.Context) {
	var err error
	messageAuth := "Вы успешно авторизировались"

	var authData model.AuthorizationData
	if err = ctx.ShouldBindJSON(&authData); err != nil {
		a.Error(ctx, http.StatusBadRequest, "некорректные данные")
		return
	}

	err = a.logic.Authorization(ctx.Request.Context(), authData.Login, authData.Password)
	if err != nil {
		a.Error(ctx, http.StatusUnauthorized, "ошибка авторизации")
		return
	}

	ctx.JSON(http.StatusOK, messageAuth)
}
