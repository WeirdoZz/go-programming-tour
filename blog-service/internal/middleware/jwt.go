package middleware

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := context.GetQuery("token"); exist {
			// 先看是否在query参数中
			token = s
		} else {
			// 再看是否放在了header中
			token = context.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(context)
			response.ToErrorResponse(ecode)
			// 如果token验证不通过就放弃接下来的操作
			context.Abort()
			return
		}
		context.Next()
	}
}
