package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/model"
	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/utils"
	_ "github.com/lib/pq"
)

type DBClient struct {
	db      *sql.DB
	siteURL string
}

func New(dbConnection string, siteURL string) (DBClient, error) {
	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", dbConnection)
	if err != nil {
		return DBClient{}, err
	}

	// Проверяем подключение к базе данных
	err = db.Ping()
	if err != nil {
		db.Close()
		return DBClient{}, err
	}

	return DBClient{
		db:      db,
		siteURL: siteURL,
	}, nil
}

func (c *DBClient) InsertRecord(ctx context.Context, data *model.InsertDataForURLs) error {
	// Проверяем и обновляем префикс в длинной ссылке
	longURL := utils.NormalizeLongURL(data.LongURL)

	// Выполняем SQL-запрос с передачей параметров longURL и shortURL
	var id int
	err := c.db.QueryRowContext(ctx, insertShortURL, longURL, data.ShortURL, data.UserID).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (c *DBClient) FindLongURL(ctx context.Context, shortURL string) (string, error) {
	var longURL string
	err := c.db.QueryRowContext(ctx, getLongURLByShort, shortURL).Scan(&longURL)
	if err != nil {
		return "", err
	}

	return longURL, nil
}

func (c *DBClient) InsertUserData(ctx context.Context, userData model.UserData) error {
	var shortID int
	shortUrl := c.siteURL + userData.ShortURL
	err := c.db.QueryRowContext(ctx, getURLIDByShort, shortUrl).Scan(&shortID)
	if err != nil {
		return err
	}
	_, err = c.db.ExecContext(
		ctx,
		insertURLInfo,
		shortID,
		userData.Browser,
		userData.BrowserVersion,
		userData.Model,
		userData.Platform,
		userData.OS,
		userData.IP,
		userData.ShortURL,
		userData.Bot,
		userData.Mobile,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *DBClient) FindInfoShortURL(ctx context.Context, shortURL string) ([]model.URLInfo, error) {
	var shortID int
	shortURL = c.siteURL + shortURL
	err := c.db.QueryRowContext(ctx, getURLIDByShort, shortURL).Scan(&shortID)
	if err != nil {
		return nil, err
	}

	rows, err := c.db.QueryContext(ctx, getURLUsageStats, shortID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urlInfos []model.URLInfo
	for rows.Next() {
		var urlInfo model.URLInfo
		err := rows.Scan(&urlInfo.ID, &urlInfo.URLID, &urlInfo.UserAgent, &urlInfo.Device, &urlInfo.OS, &urlInfo.IPAddress)
		if err != nil {
			return nil, err
		}
		urlInfos = append(urlInfos, urlInfo)
	}

	return urlInfos, nil
}

func (c DBClient) FindDataShortURL(ctx context.Context, longURL string) (string, error) {

	var existingShortURL string
	err := c.db.QueryRowContext(ctx, getShortURLByLong, longURL).Scan(&existingShortURL)
	if err == nil {
		return existingShortURL, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	return "", nil
}
