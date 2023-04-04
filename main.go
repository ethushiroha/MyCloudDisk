package main

import (
	"MyCloudDisk/global"
	"MyCloudDisk/router"
	"MyCloudDisk/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化全局资源
	storage.Setup()
	defer storage.Close()
	// 创建Router
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run(global.VP.GetString("server.backend_url"))
	if err != nil {
		return
	}
}
