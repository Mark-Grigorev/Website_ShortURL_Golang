package main

func findshortUrl(shortID string) (string, error) {

	db, err := connectToPostgreSQL()
	defer db.Close()
	shortURL := "http://localhost:8080/" + shortID
	longUrl, err := findLongURL(db, shortURL)
	if err != nil {
		return "", err
	}
	return longUrl, nil
}
