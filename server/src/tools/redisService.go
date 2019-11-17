package tools

import "github.com/go-redis/redis/v7"

func CreateRedisService(Addr string,passwd string)  * redis.Client {
	client := redis.NewClient(&redis.Options{
		DB:       0,  // use default DB
		Password:passwd,
		Addr:Addr,
	})
	return  client
}

var redisClient * redis.Client = nil

func init() {
	redisClient = CreateRedisService("localhost:6379","")
}

func GetRedisClient() * redis.Client {
	return  redisClient
}