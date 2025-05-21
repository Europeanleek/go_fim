package core

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func InitRedis(addr, pwd string, db int) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       db,  // use default DB
		PoolSize: 100, // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis连接成功")
	return
}
