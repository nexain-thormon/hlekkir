package client

import (
	"github.com/redis/go-redis/v9"
)

func Redis(redisURL string) *redis.Client {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}
