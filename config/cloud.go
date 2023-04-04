package config

import "MyCloudDisk/global"

func InitCloudDisk() string {
	vp := global.VP
	return vp.GetString("disk.home")
}
