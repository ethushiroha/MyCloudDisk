package middleware

import (
	"MyCloudDisk/service"
	"MyCloudDisk/utils/ApiUtils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func parseJwt(authHeader string) (*jwt.Token, error) {
	// 解析 token
	return jwt.ParseWithClaims(authHeader, &ApiUtils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method: " + token.Header["alg"].(string))
		}
		return ApiUtils.GetSecretKey()
	})
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			ApiUtils.UnauthorizedResponse(c)
			return
		}
		authHeader = strings.TrimPrefix(authHeader, "Bearer ")

		// 如果 token 在缓存里不存在，则登录失败
		if !service.GetATokenFromRedisCache(authHeader) {
			ApiUtils.UnauthorizedResponse(c)
			return
		}

		token, err := parseJwt(authHeader)
		// 如果解析成功，将解析后的信息保存到上下文中
		if err == nil {
			if claims, ok := token.Claims.(*ApiUtils.Claims); ok && token.Valid {
				// 权限不对
				//if Global.RoleGuest < claims.Role {
				//	ApiUtils.UnauthorizedResponse(c)
				//	return
				//}
				//c.Set("id", claims.ID)
				c.Set("UserID", claims.ID)
				c.Next()
				return
			}
		}

		ApiUtils.UnauthorizedResponse(c)
		return
	}
}
