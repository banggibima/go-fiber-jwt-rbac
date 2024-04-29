package redis

import (
	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/redis/go-redis/v9"
)

func New(config *config.Config) (*redis.Client, error) {
	addr := config.Redis.Addr
	password := config.Redis.Password
	database := config.Redis.Database

	options := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       database,
	}

	client := redis.NewClient(options)

	if err := Connect(client); err != nil {
		return nil, err
	}

	return client, nil
}
