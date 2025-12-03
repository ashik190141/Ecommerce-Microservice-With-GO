package radis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitializeRedisClient(redisURL string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	if(rdb == nil){
		panic("Failed to connect to Redis")
	}else{
		fmt.Println("Connected to Redis successfully")
	}

	return rdb
}