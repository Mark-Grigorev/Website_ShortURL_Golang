package main

import (
	"database/sql"

	"github.com/teris-io/shortid"
)

func FindshortUrl(shortID string) (string, error) {

	db, err := ConnectToPostgreSQL()
	defer db.Close()
	shortURL := "http://localhost:8080/" + shortID
	longUrl, err := FindLongURL(db, shortURL)
	if err != nil {
		return "", err
	}
	return longUrl, nil
}

func GenerateshortURL(longURL string) (string, error) {
	localhost := "http://localhost:8080/"

	// Проверяем, есть ли уже такая запись в базе данных
	db, err := ConnectToPostgreSQL()
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Запрос для проверки наличия записи с таким longURL
	query := "SELECT shorturl FROM urls_users WHERE longurl = $1"
	var existingShortURL string
	longURL = NormalizeLongURL(longURL)
	err = db.QueryRow(query, longURL).Scan(&existingShortURL)
	if err == nil {
		// Запись уже существует, возвращаем существующий shortURL
		return existingShortURL, nil
	} else if err != sql.ErrNoRows {
		// Произошла ошибка при выполнении запроса
		return "", err
	}

	// Генерируем новый shortURL, так как записи с таким longURL нет в базе данных
	id, err := shortid.Generate()
	if err != nil {
		return "", err
	}

	shortURL := localhost + id

	// Вызываем функцию для создания новой записи с заданными значениями полей longURL и shortURL
	err = InsertRecord(db, longURL, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func StatusShortUrl(shortID string) ([]URLInfo, error) {
	db, err := ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	info, err := FindInfoShortURL(db, shortID)
	if err != nil {
		return nil, err
	}

	return info, nil
}
