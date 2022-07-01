package middleware

import (
	"github.com/gin-gonic/gin"
	"go-depot/pkg/app"
	"go-depot/pkg/errcode"
	"go-depot/pkg/limiter"
)

/**
入参应该为 LimiterIface 接口类型，这样子的话只要符合该接口类型的具体限流器实现都可以传入并使用
另外比较重要的就是 TakeAvailable 方法，它会占用存储桶中立即可用的令牌的数量，返回值为删除的令牌数，如果没有可用的令牌，将会返回 0，也就是已经超出配额了
因此这时候我们将返回 errcode.TooManyRequest 状态告诉客户端需要减缓并控制请求速度。
*/
func RateLimiter(l limiter.LimiterInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
