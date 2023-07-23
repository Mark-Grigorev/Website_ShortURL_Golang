package helpers

import "strings"

// Функция для проверки и обновления префикса "https://" в длинной ссылке
func NormalizeLongURL(longURL string) string {
	if !strings.HasPrefix(longURL, "http://") && !strings.HasPrefix(longURL, "https://") {
		longURL = "https://" + longURL
	}
	return longURL
}
