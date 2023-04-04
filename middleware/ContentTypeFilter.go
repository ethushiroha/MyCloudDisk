package middleware

import (
	"MyCloudDisk/utils/ApiUtils"
	"github.com/gin-gonic/gin"
)

func ContentTypeFilter(contentType string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Method == "GET" {
			return
		}
		thisType := context.Request.Header.Get("Content-Type")
		if thisType != contentType {
			ApiUtils.BadRequestResponse(context, "Content-Type must be "+contentType)
		}
		context.Next()
	}
}
