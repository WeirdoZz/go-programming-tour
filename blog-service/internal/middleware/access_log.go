package middleware

import (
	"blog-service/global"
	"blog-service/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 向body中写入p
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

// AccessLog 获取日志的中间件
func AccessLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: context.Writer,
		}
		context.Writer = bodyWriter

		beginTime := time.Now().Unix()
		context.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  context.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		s := "access log: method:%s,status_code:%d," + "begin_time:%d,end_time:%d"
		global.Logger.WithFields(fields).Infof(context, s,
			context.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime)
	}
}
