package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis faield:%v\n", err)
		return
	}
	fmt.Println("connect redis success！")
}
