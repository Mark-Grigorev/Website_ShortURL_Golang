package db

const (
	// Запрос для вставки новой короткой ссылки.
	insertShortURL = "INSERT INTO urls_users (longurl, shorturl, user_id) VALUES ($1, $2, $3) RETURNING id"
	// Запрос для получения оригинального URL по короткому.
	getLongURLByShort = "SELECT longurl FROM urls_users WHERE shorturl = $1"
	// Запрос для получения ID записи по короткому URL.
	getURLIDByShort = "SELECT id FROM urls_users WHERE shorturl = $1"
	// Запрос для вставки информации о переходах (для статистики).
	insertURLInfo = `INSERT INTO urls_info 
	(url_id, user_browser, user_browser_version, user_model, user_platform, user_os, user_ip, user_short_url, user_bot, user_mobile) 
	VALUES ($1, $2, $3, $4, $5)`
	// Запрос для получения статистики по конкретному URL.
	getURLUsageStats = "SELECT * FROM urls_info WHERE url_id = $1"
	// Запрос для поиска существующего короткого URL по длинному.
	getShortURLByLong = "SELECT shorturl FROM urls_users WHERE longurl = $1"
)
