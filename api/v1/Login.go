package v1

import (
	"MyCloudDisk/model/ApiModels"
	"MyCloudDisk/service"
	"MyCloudDisk/utils/ApiUtils"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var body map[string]interface{}
	err := context.ShouldBindJSON(&body)
	if err != nil {
		ApiUtils.ErrorResponse(context, err)
		return
	}
	var username, password string

	if u, ok := body["Username"]; !ok {
		ApiUtils.BadRequestResponse(context, "Param error: Username is required")
		return
	} else {
		username = u.(string)
	}
	if p, ok := body["Password"]; !ok {
		ApiUtils.BadRequestResponse(context, "Param error: Password is required")
		return
	} else {
		password = p.(string)
	}

	user := service.GetUser(username, password)
	if user == nil {
		ApiUtils.UnauthorizedResponse(context)
		context.Abort()
		return
	}

	token, err := ApiUtils.GenerateToken(uint(user.UserID), username)
	if err != nil {
		ApiUtils.ErrorResponse(context, err)
		return
	}
	data := map[string]interface{}{
		"username": username,
		"token":    token,
	}
	service.SetTokenToRedisCache(token)

	resp := &ApiModels.LoginRespBody{}
	resp.SetData(data).
		SetStatus("success").
		SetMessage("登录成功").
		Build()

	//context.Set("UserID", user.UserID)
	ApiUtils.SuccessResponse(context, resp)
}
