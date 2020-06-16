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

	// zset
	key := "rank"
	items := []*redis.Z{
		&redis.Z{Score: 90, Member: "PHP"},
		&redis.Z{Score: 91, Member: "Java"},
		&redis.Z{Score: 92, Member: "Python"},
		&redis.Z{Score: 93, Member: "Golang"},
	}

	// 把元素追加到key中
	redisdb.ZAdd(key, items...)

	// 给Golang + 10分
	newScore, err := redisdb.ZIncrBy(key, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed,err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := redisdb.ZRangeWithScores(key, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed,err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisdb.ZRangeByScoreWithScores(key, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed,err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
