package logic

import (
	"context"
	"fmt"

	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/db"
	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/model"
	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/utils"
	"github.com/gin-gonic/gin"
	userAgent "github.com/mssola/user_agent"
	"github.com/teris-io/shortid"
)

type Logic struct {
	cfg *model.Config
	db  *db.DBClient
}

func New(cfg *model.Config, db *db.DBClient) *Logic {
	return &Logic{
		cfg: cfg,
		db:  db,
	}
}

func (l *Logic) GetUserData(ctx *gin.Context, shortUrl string) error {
	logPrefix := "[GetUserData]"

	agent := ctx.Request.UserAgent()
	ua := userAgent.New(agent)
	browser, browserVers := ua.Browser()
	err := l.db.InsertUserData(
		ctx.Request.Context(),
		model.UserData{
			Browser:        browser,
			BrowserVersion: browserVers,
			OS:             ua.OS(),
			Platform:       ua.Platform(),
			Model:          ua.Model(),
			IP:             ctx.ClientIP(),
			ShortURL:       shortUrl,
			Bot:            ua.Bot(),
			Mobile:         ua.Mobile(),
		})
	if err != nil {
		return fmt.Errorf("%s ошибка при записи данных в бд %w", logPrefix, err)
	}
	return nil
}

func (l *Logic) FindLongUrl(ctx context.Context, shortID string) (string, error) {
	logPrefix := "[FindLongUrl]"
	shortURL := l.cfg.AppConfig.ThisURL + shortID

	longUrl, err := l.db.FindLongURL(ctx, shortURL)
	if err != nil {
		return "", fmt.Errorf("%s ошибка при поиске короткой ссылки %w", logPrefix, err)
	}
	return longUrl, nil
}

func (l *Logic) GetOrCreateShortURL(ctx context.Context, longURL string) (string, error) {
	logPrefix := "[GetOrCreateShortURL]"
	var shortURL string

	shortURL, err := l.db.FindDataShortURL(
		ctx,
		utils.NormalizeLongURL(longURL),
	)
	if err != nil {
		return "", fmt.Errorf("%s ошибка при получении короткой ссылки %w", logPrefix, err)
	}
	if shortURL == "" {
		// Генерируем новый shortURL, так как записи с таким longURL нет в базе данных
		id, err := shortid.Generate()
		if err != nil {
			return "", fmt.Errorf("%s ошибка при создании уникального ShortID %w", logPrefix, err)
		}

		shortURL := l.cfg.AppConfig.ThisURL + id

		// Вызываем функцию для создания новой записи с заданными значениями полей longURL и shortURL
		err = l.db.InsertRecord(ctx, longURL, shortURL)
		if err != nil {
			return "", fmt.Errorf("%s ошибка при записи в бд %w", logPrefix, err)
		}
	}

	return shortURL, nil
}

func (l *Logic) StatusShortUrl(ctx context.Context, shortID string) ([]model.URLInfo, error) {
	return l.db.FindInfoShortURL(ctx, shortID)
}
