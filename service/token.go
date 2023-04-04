package service

import (
	"MyCloudDisk/global"
	"github.com/go-redis/redis"
)

func GetATokenFromRedisCache(token string) bool {
	_, err := global.RedisDB.Get(token).Result()
	if err == redis.Nil || err != nil {
		return false
	} else {
		return true
	}
}

func SetTokenToRedisCache(token string) bool {
	timeout := global.JwtSecret.Timeout
	_, err := global.RedisDB.Set(token, true, timeout).Result()
	return err == nil
}

func RemoveTokenFromRedisCache(token string) bool {
	_, err := global.RedisDB.Del(token).Result()
	return err == nil
}
