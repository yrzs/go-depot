package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-depot/global"
	"go-depot/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

/**
实际上在写入流时，调用的是 http.ResponseWriter
实现我们特定的 Write 方法就可以解决无法直接取到方法响应主体的问题了
*/
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

/**
类似nginx access_log
*/
func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWrite := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = bodyWrite

		//调用方法的开始时间，调用方法结束的结束时间。
		beginTime := time.Now().Unix()
		ctx.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  ctx.Request.PostForm.Encode(), //当前的请求参数。
			"response": bodyWrite.body.String(),       //当前的请求结果响应主体。
		}
		// write log
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			ctx.Request.Method, //当前的调用方法。
			bodyWrite.Status(), //当前的响应结果状态码。
			beginTime,
			endTime,
		)
	}
}
