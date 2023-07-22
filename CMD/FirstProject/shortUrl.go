package main

import (
	"strings"

	"github.com/teris-io/shortid"
)

func generateShortURL(longURL string) (string, error) {
	id, err := shortid.Generate()
	if err != nil {
		return "", err
	}
	localhost := "http://localhost:8080/"
	shortURL := localhost + id

	db, err := connectToPostgreSQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Вызываем функцию для создания новой записи с заданными значениями полей longURL и shortURL
	err = insertRecord(db, longURL, shortURL)
	if err != nil {
		panic(err)
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
