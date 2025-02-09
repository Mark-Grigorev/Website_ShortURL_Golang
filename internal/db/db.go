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

func (c *DBClient) InsertRecord(ctx context.Context, longURL, shortURL string) error {
	// Проверяем и обновляем префикс в длинной ссылке
	longURL = utils.NormalizeLongURL(longURL)

	// Выполняем SQL-запрос с передачей параметров longURL и shortURL
	var id int
	err := c.db.QueryRowContext(ctx, insertShortURL, longURL, shortURL).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (c *DBClient) FindLongURL(ctx context.Context, shortURL string) (string, error) {
	// Выполняем SQL-запрос с передачей параметра shortURL и получаем результат
	var longURL string
	err := c.db.QueryRowContext(ctx, getLongURLByShort, shortURL).Scan(&longURL)
	if err != nil {
		return "", err
	}
	//Удаляем часть Url, для корректной работы

	return longURL, nil
}

func (c *DBClient) InsertUserData(ctx context.Context, userData model.UserData) error {
	// Запрос для получения ID короткой ссылки из таблицы urls_users

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

// TODO - Вынести регистрацию/авторизацию в отдельный сервис
func (c *DBClient) RegistrationUserDB(ctx context.Context, login string, password string, name string) error {
	_, err := c.db.ExecContext(ctx, insertUser, login, password, name)
	if err != nil {
		return err
	}
	return nil
}

func (c *DBClient) AuthorizationUserDB(ctx context.Context, login string, password string) error {
	var count int

	// Выполняем запрос с переданными параметрами логина и пароля
	err := c.db.QueryRow(checkUserCredentials, login, password).Scan(&count)
	if err != nil {
		return err
	}

	// Если count > 0, значит есть совпадение логина и пароля в базе данных
	if count > 0 {
		return nil
	}

	// Если нет совпадений, возвращаем ошибку
	return errors.New("пользователь с таким логином и паролем не найден")
}
