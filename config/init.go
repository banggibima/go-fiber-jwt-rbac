package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App      App
	HTTP     HTTP
	Postgres Postgres
	Redis    Redis
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

type Redis struct {
	Addr     string
	Password string
	Database int
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
		Redis: Redis{
			Addr:     viper.GetString("REDIS_ADDR"),
			Password: viper.GetString("REDIS_PASSWORD"),
			Database: viper.GetInt("REDIS_DATABASE"),
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
