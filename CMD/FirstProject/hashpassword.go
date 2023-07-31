package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func HelperHashPass(password string) (string, error) {
	// Создаем объект хэша SHA256
	hasher := sha256.New()

	// Записываем пароль в хэшер
	hasher.Write([]byte(password))

	// Получаем хэш в бинарном формате
	hashedPassword := hasher.Sum(nil)

	// Преобразуем хэш в строку в формате hex
	hashedPasswordStr := hex.EncodeToString(hashedPassword)

	return hashedPasswordStr, nil
}
