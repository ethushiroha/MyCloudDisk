package storage

import (
	"MyCloudDisk/config"
	"MyCloudDisk/global"
)

func Setup() {
	global.VP = config.InitViper()
	global.DB = config.InitMySQL()
	global.RedisDB = config.InitRedis()
	global.JwtSecret = config.InitJwt()
	global.CloudDiskBasicDir = config.InitCloudDisk()
}

func Close() {
	global.DB.Close()
	global.RedisDB.Close()
}
