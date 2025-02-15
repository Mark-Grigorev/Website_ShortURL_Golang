package model

type Config struct {
	DBConfig   DBConfig
	AppConfig  AppConfig
	AuthConfig AuthConfig
}

type AppConfig struct {
	Host    string
	ThisURL string
	Debug   string
}
type DBConfig struct {
	DBConnection string
}

type AuthConfig struct {
	Host string
}
