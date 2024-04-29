package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Connect(rdb *redis.Client) error {
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return err
	}

	return nil
}
