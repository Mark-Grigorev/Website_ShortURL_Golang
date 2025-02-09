package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// Функция для проверки и обновления префикса "https://" в длинной ссылке
func NormalizeLongURL(longURL string) string {
	if !strings.HasPrefix(longURL, "http://") && !strings.HasPrefix(longURL, "https://") {
		longURL = "https://" + longURL
	}
	return longURL
}

func HashPass(password string) (string, error) {
	// Создаем объект хэша SHA256
	hasher := sha256.New()

	// Записываем пароль в хэшер
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return "", err
	}
	// Получаем хэш в бинарном формате
	hashedPassword := hasher.Sum(nil)

	// Преобразуем хэш в строку в формате hex
	hashedPasswordStr := hex.EncodeToString(hashedPassword)

	return hashedPasswordStr, nil
}
