package main

import (
	"database/sql"
	"strings"

	"github.com/teris-io/shortid"
)

func generateShortURL(longURL string) (string, error) {
	localhost := "http://localhost:8080/"

	// Проверяем, есть ли уже такая запись в базе данных
	db, err := connectToPostgreSQL()
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Запрос для проверки наличия записи с таким longURL
	query := "SELECT shorturl FROM urls_users WHERE longurl = $1"
	var existingShortURL string
	longURL = normalizeLongURL(longURL)
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
	err = insertRecord(db, longURL, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

// Функция для проверки и обновления префикса "https://" в длинной ссылке
func normalizeLongURL(longURL string) string {
	if !strings.HasPrefix(longURL, "http://") && !strings.HasPrefix(longURL, "https://") {
		longURL = "https://" + longURL
	}
	return longURL
}
