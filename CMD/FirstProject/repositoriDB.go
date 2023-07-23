package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func connectToPostgreSQL() (*sql.DB, error) {
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

func insertRecord(db *sql.DB, longURL, shortURL string) error {
	// Проверяем и обновляем префикс в длинной ссылке
	longURL = normalizeLongURL(longURL)

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

func findLongURL(db *sql.DB, shortURL string) (string, error) {
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

func insertUserData(db *sql.DB, shortUrl string, userData *UserData) error {
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
