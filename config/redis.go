package config

import (
	"MyCloudDisk/global"
	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	vp := global.VP
	Host := vp.GetString("redis.Host")
	Password := vp.GetString("redis.Password")

	RedisDB := redis.NewClient(&redis.Options{
		Addr:     Host,
		Password: Password, // no password set
		DB:       0,        // use default MySqlDB
	})
	_, err := RedisDB.Ping().Result()
	if err != nil {
		panic("初始化失败：redis 连接失败")
	}
	return RedisDB
}
