
package utils


import "github.com/redis/go-redis/v9"

func Connect(addr string, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	defer rdb.Close()

	return rdb
}
