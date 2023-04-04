package ApiUtils

import (
	"MyCloudDisk/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID   uint   `json:"id"`
	Role int    `json:"role"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func GetSecretKey() ([]byte, error) {
	return global.JwtSecret.Secret, nil
}

func GenerateToken(id uint, name string) (string, error) {
	// 定义 JWT 中包含的信息
	claims := &Claims{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * global.JwtSecret.Timeout).Unix(), // 过期时间为 1 小时
			Issuer:    "MyBurp",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key, err := GetSecretKey()
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}
