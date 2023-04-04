package config

import (
	"MyCloudDisk/global"
	"time"
)

func InitJwt() *global.MyJwt {
	vp := global.VP
	secret := vp.GetString("jwt.secret")
	timeout := vp.GetInt("jwt.timeout")
	return &global.MyJwt{
		Secret:  []byte(secret),
		Timeout: time.Duration(timeout) * time.Second,
	}
}
