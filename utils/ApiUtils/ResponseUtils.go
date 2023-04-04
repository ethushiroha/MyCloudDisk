package ApiUtils

import (
	"MyCloudDisk/model/ApiModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(c *gin.Context, err error) {
	resp := ApiModels.BasicRespBody{
		Status:  "Error",
		Message: err.Error(),
	}
	c.JSON(http.StatusInternalServerError, resp)
	// todo: 记录 error
	c.Abort()
}

func UnauthorizedResponse(c *gin.Context) {
	resp := ApiModels.BasicRespBody{
		Status:  "failed",
		Message: "未登录 或 token 过期",
	}
	c.JSON(http.StatusUnauthorized, resp)
	c.Abort()
}

func SuccessResponse(c *gin.Context, data ApiModels.RespBody) {
	c.JSON(http.StatusOK, data)
}

func FailureResponse(c *gin.Context, message string) {
	resp := ApiModels.BasicRespBody{
		Status:  "failed",
		Message: message,
	}
	c.JSON(http.StatusOK, resp)
}

func DirectResponse(c *gin.Context, uri string) {
	c.Redirect(http.StatusFound, uri)
}

func BadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "failure",
		"message": message,
	})
	c.Abort()
}
