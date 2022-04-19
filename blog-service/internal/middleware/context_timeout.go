package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// ContextTimeout 设置当前context的超时时间
func ContextTimeout(t time.Duration) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(ctx.Request.Context(), t)
		defer cancel()
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}
