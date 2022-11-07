package cache

import (
	"github.com/go-redis/redis"
	"log"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	//TODO 从本地文件加载
}

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		//Password:
		DB: 1,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	RedisClient = client
}

func LoadRedisData() {
	RedisDb = ""
	RedisAddr = "localhost:6379"
	RedisPw = ""
	RedisDbName = ""
}
