package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-depot/global"
	"go-depot/pkg/app"
	"go-depot/pkg/errcode"
	"go-depot/pkg/wechat"
)

/**
gin 本身已经自带了一个 Recovery 中间件:gin.Recovery()
但是在项目中需要针对我们的公司内部情况或生态圈定制 Recovery 中间件，
确保异常在被正常捕抓之余，要及时的被识别和处理，因此自定义一个 Recovery 中间件是非常有必要的
*/
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				// send wechat webhook
				go func(c *gin.Context) { // 异步发送
					if err := recover(); err != nil {
						global.Logger.Errorf(c, "wechat.webhook.SendTextMsg err: %v", err)
					}
					wechat.SendHookTextMsg(
						fmt.Sprintf(
							"%s -- check the app logs view stack info or see http://%s:16686",
							global.Logger.Errorf4Webhook(c, "panic recover err: %v", err),
							global.AppSetting.OpenTracing.AgentHost,
						),
					)
				}(c)
				c.Abort()
			}
		}()
		c.Next()
	}
}

/**
永远不要启动一个你无法控制它退出，或者你无法知道它何时推出的 goroutine
启动 goroutine 时请加上 panic recovery 机制，避免服务直接不可用
造成 goroutine 泄漏的主要原因就是 goroutine 中造成了阻塞，并且没有外部手段控制它退出
尽量避免在请求中直接启动 goroutine 来处理问题,而应该通过启动 worker 来进行消费，这样可以避免由于请求量过大，而导致大量创建 goroutine 从而导致 oom，当然如果请求量本身非常小，那当我没说
*/
//func Go(f func()) {
//	go func() {
//		defer func() {
//			if err := recover(); err != nil {
//				global.Logger.Errorf(c, "wechat.webhook.SendTextMsg err: %v", err)
//			}
//		}()
//		f()
//	}()
//}
