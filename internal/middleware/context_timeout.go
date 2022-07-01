package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

/**
统一的在应用程序中针对所有请求都进行一个最基本的超时时间控制
*/
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
