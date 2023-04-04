package middleware

import (
	"MyCloudDisk/model/ApiModels"
	"MyCloudDisk/utils/ApiUtils"
	"github.com/gin-gonic/gin"
)

func RoleFilter(role int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exists := ctx.Get("role")
		if !exists {
			ApiUtils.UnauthorizedResponse(ctx)
			return
		} else {
			if value.(int) > role {
				resp := &ApiModels.BasicRespBody{}
				resp.SetStatus("failure").
					SetMessage("权限不足").
					Build()
				ApiUtils.SuccessResponse(ctx, resp)
				ctx.Abort()
				return
			} else {
				ctx.Next()
			}
		}
	}
}
