package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App      App
	HTTP     HTTP
	Postgres Postgres
	JWT      JWT
}

type App struct {
	Name    string
	Version string
}

type HTTP struct {
	Port int
}

type Postgres struct {
	Driver     string
	Connection string
}

type JWT struct {
	AccessSecret  string
	RefreshSecret string
	AccessExpiry  int
	RefreshExpiry int
	Audience      string
	Issuer        string
}

func Init(viper *viper.Viper) (*Config, error) {
	config := &Config{
		App: App{
			Name:    viper.GetString("APP_NAME"),
			Version: viper.GetString("APP_VERSION"),
		},
		HTTP: HTTP{
			Port: viper.GetInt("HTTP_PORT"),
		},
		Postgres: Postgres{
			Driver:     viper.GetString("POSTGRES_DRIVER"),
			Connection: viper.GetString("POSTGRES_CONNECTION"),
		},
		JWT: JWT{
			AccessSecret:  viper.GetString("JWT_ACCESS_SECRET"),
			RefreshSecret: viper.GetString("JWT_REFRESH_SECRET"),
			AccessExpiry:  viper.GetInt("JWT_ACCESS_EXPIRY"),
			RefreshExpiry: viper.GetInt("JWT_REFRESH_EXPIRY"),
			Audience:      viper.GetString("JWT_AUDIENCE"),
			Issuer:        viper.GetString("JWT_ISSUER"),
		},
	}

	return config, nil
}
