package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type Config struct {
	Address  string
	Password string
}

func New(ctx context.Context, config *Config) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
	}), nil
}
