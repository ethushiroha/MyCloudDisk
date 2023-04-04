package global

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"time"
)

type MyJwt struct {
	Secret  []byte
	Timeout time.Duration
}

var (
	DB                *sqlx.DB      // MySQL实例
	VP                *viper.Viper  // Viper实例
	RedisDB           *redis.Client // redis
	JwtSecret         *MyJwt
	CloudDiskBasicDir string
)

const (
	RoleAdmin = iota
	RoleUser
	RoleGuest
)

//var ContentTypeMapping = map[int]string
