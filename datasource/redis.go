package datasource

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisDatabase struct {
	RedisDBConnector *redis.Client
}

func RedisConnect() *RedisDatabase {
	redisUrl := os.Getenv("redis")
	fmt.Println(redisUrl)
	opt, redisEr := redis.ParseURL(redisUrl)

	if redisEr != nil {
		panic(redisEr)
	}
	rdb := redis.NewClient(opt)
	return &RedisDatabase{
		RedisDBConnector: rdb,
	}
}
