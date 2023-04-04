package router

import (
	v1 "MyCloudDisk/api/v1"
	"MyCloudDisk/controller"
	"MyCloudDisk/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", controller.Login)

	home := r.Group("/home")
	{
		//r.GET("/index", controller.Index)
		home.GET("/index", controller.Index)
		home.GET("/files", controller.Files)
		home.GET("/upload", controller.Upload)
	}

	apiV1 := r.Group("/api/v1")
	//apiV1.Use(middleware.AuthorizeJWT())
	apiV1.POST("/login", v1.Login)
	//apiV1.Use(middleware.ContentTypeFilter("application/json"))
	//apiV1.POST("/home/upload", v1.Upload)

	apiV1Home := apiV1.Group("/home")
	apiV1Home.Use(middleware.AuthorizeJWT())
	{
		apiV1Home.POST("/upload", v1.Upload)
		apiV1Home.POST("/index", v1.Index)
		apiV1Home.POST("/files", v1.Files)
		apiV1Home.POST("/addFolder", v1.AddFolder)
		apiV1Home.POST("/download", v1.DownloadFile)
	}

	r.LoadHTMLGlob("./view/*")
	r.Static("/static", "./static")

}
