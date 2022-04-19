package middleware

import (
	"blog-service/global"
	"blog-service/pkg/Email"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := Email.NewEmail(&Email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err:%v"
				global.Logger.WithCallersFrames().Errorf(context, s, err)

				err := defaultMailer.SendMail(global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err))
				if err != nil {
					global.Logger.Panicf(context, "mail.SendMail err:%v", err)
				}
				app.NewResponse(context).ToErrorResponse(errcode.ServerError)
				context.Abort()
			}
		}()
		context.Next()
	}
}
