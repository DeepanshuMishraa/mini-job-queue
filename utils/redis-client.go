package utils

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func Connect(url string) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opts)
	log.Println("Connected To Redis")
	return rdb, nil
}
