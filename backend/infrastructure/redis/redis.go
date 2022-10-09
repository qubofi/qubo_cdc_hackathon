package redis

import (
	"os"

	"github.com/go-redis/redis/v7"
)

var Client *redis.Client

func SetupRedis() {
	dsn := os.Getenv("REDIS_DSN")
	if dsn == "" {
		dsn = "localhost:6379"
	}

	Client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}

	println("Successfully connected to redis!")
}
