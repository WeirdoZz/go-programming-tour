package app

import (
	"blog-service/global"
	"blog-service/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

// GetJWTSecret 获取项目的JWT Secret
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 生成token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)

	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	// 生成header和payload
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成signature
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// ParseToken 解析token字符串，并且返回其中蕴含的claims
func ParseToken(token string) (*Claims, error) {
	// 将经过生成的token字符串进行解析，生成token结构体
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenClaims != nil {
		// 看是否能从token结构体中获取到信息
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

