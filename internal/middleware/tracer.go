package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"go-depot/global"
)

/**
jaeger链路追踪中间件
*/
func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var newCtx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{
					Key:   string(ext.Component),
					Value: "HTTP",
				},
			)
		}
		defer span.Finish()

		var traceID, spanID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			traceID = spanContext.(jaeger.SpanContext).TraceID().String() // 对spanContext的 jaeger.SpanContext 做断言，获取SpanID 和traceID
			spanID = spanContext.(jaeger.SpanContext).SpanID().String()
		}
		c.Set("X-Trace-ID", traceID) // 带上本次请求的链路属性
		c.Set("X-Span-ID", spanID)

		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
