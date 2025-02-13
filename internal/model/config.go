package model

type Config struct {
	DBConfig  DBConfig
	AppConfig AppConfig
}

type AppConfig struct {
	Host    string
	ThisURL string
	Debug   string
}
type DBConfig struct {
	DBConnection string
}
