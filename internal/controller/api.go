package controller

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/logic"
	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/model"
	"github.com/gin-gonic/gin"
)

type App struct {
	config *model.Config
	router *gin.Engine
	logic  *logic.Logic
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func New(cfg *model.Config, logic *logic.Logic) *App {
	return &App{
		router: setupRouter(),
		config: cfg,
		logic:  logic,
	}

}
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
	return router
}

func (a *App) setV1Routes() {
	v1 := a.router.Group("/v1")
	a.setURLShorterRoutes(v1.Group("/urlshort"))
	a.setAuthRoutes(v1.Group("/auth"))

}

func (a *App) Start() {
	a.setV1Routes()
	server := &http.Server{
		Addr:    ":" + a.config.Host,
		Handler: a.router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		log.Printf("Сервер запускается на порту - %s", a.config.Host)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Ошибка при запуске сервера - %s", err.Error())
		}
	}()

	<-ctx.Done()
	log.Printf("Выключение сервера")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Не удалось завершить работу сервера - %s \n", err.Error())
	}
	log.Println("Сервер корректно остановлен")
}

func (a *App) Error(ctx *gin.Context, status int, message string) {
	ctx.JSON(
		status,
		ErrorResponse{
			Status:  status,
			Message: message,
		})
}
