package core

import (
	"github.com/go-redis/redis"
	"log"
)

// 声明一个全局的rdb变量
var RDB *redis.Client

// 初始化连接
func initRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		log.Fatalf(" [ERROR] redis connect :%s err:%v\n", RDB, err)
	}

}
