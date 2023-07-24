package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UserData struct {
	UserAgent string
	Device    string
	OS        string
	IP        string
}

func ConnectToPostgreSQL() (*sql.DB, error) {
	//Строка подключения
	dbinfo := "user=postgres password=123 dbname=ShortUrl host=localhost port=5432 sslmode=disable"

	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	// Проверяем подключение к базе данных
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	//fmt.Println("Успешное подключение к базе данных PostgreSQL!")

	return db, nil
}

func InsertRecord(db *sql.DB, longURL, shortURL string) error {
	// Проверяем и обновляем префикс в длинной ссылке
	longURL = NormalizeLongURL(longURL)

	// Запрос для вставки новой записи в таблицу
	query := "INSERT INTO urls_users (longurl, shorturl) VALUES ($1, $2) RETURNING id"

	// Выполняем SQL-запрос с передачей параметров longURL и shortURL
	var id int
	err := db.QueryRow(query, longURL, shortURL).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func FindLongURL(db *sql.DB, shortURL string) (string, error) {
	// Запрос для поиска записи по shorturl
	query := "SELECT longurl FROM urls_users WHERE shorturl = $1"

	// Выполняем SQL-запрос с передачей параметра shortURL и получаем результат
	var longURL string
	err := db.QueryRow(query, shortURL).Scan(&longURL)
	if err != nil {
		return "", err
	}
	//Удаляем часть Url, для корректной работы

	return longURL, nil
}

func InsertUserData(db *sql.DB, shortUrl string, UserAgent string, Device string, OS string, IP string) error {
	userData := &UserData{
		UserAgent: UserAgent,
		Device:    Device,
		OS:        OS,
		IP:        IP,
	}
	// Запрос для получения ID короткой ссылки из таблицы urls_users
	shortIDQuery := "SELECT id FROM urls_users WHERE shorturl = $1"
	var shortID int
	shortUrl = "http://localhost:8080/" + shortUrl
	err := db.QueryRow(shortIDQuery, shortUrl).Scan(&shortID)
	if err != nil {
		return err
	}

	// Запрос для вставки данных пользователя в таблицу urls_info
	query := "INSERT INTO urls_info (url_id, user_agent, device, os, ip_address) VALUES ($1, $2, $3, $4, $5)"
	_, err = db.Exec(query, shortID, userData.UserAgent, userData.Device, userData.OS, userData.IP)
	if err != nil {
		return err
	}

	return nil
}

type URLInfo struct {
	ID        int
	URLID     int
	UserAgent string
	Device    string
	OS        string
	IPAddress string
}

func FindInfoShortURL(db *sql.DB, shortURL string) ([]URLInfo, error) {
	shortIDQuery := "SELECT id FROM urls_users WHERE shorturl = $1"
	var shortID int
	shortURL = "http://localhost:8080/" + shortURL
	err := db.QueryRow(shortIDQuery, shortURL).Scan(&shortID)
	if err != nil {
		return nil, err
	}

	shortURLStatus := "SELECT * FROM urls_info WHERE url_id = $1"
	rows, err := db.Query(shortURLStatus, shortID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urlInfos []URLInfo
	for rows.Next() {
		var urlInfo URLInfo
		err := rows.Scan(&urlInfo.ID, &urlInfo.URLID, &urlInfo.UserAgent, &urlInfo.Device, &urlInfo.OS, &urlInfo.IPAddress)
		if err != nil {
			return nil, err
		}
		urlInfos = append(urlInfos, urlInfo)
	}

	return urlInfos, nil
}
