package Redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client = nil

func GetInstance() *redis.Client {
	if rdb == nil {
		rdb = connect()
	}
	return rdb
}

func connect() *redis.Client {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := rdb.Set(ctx, "key", "value", 0).Err(); err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	return rdb
}
