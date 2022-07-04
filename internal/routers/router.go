package routers

import (
	"github.com/gin-gonic/gin"
	"go-depot/global"
	"go-depot/internal/middleware"
	v1 "go-depot/internal/routers/api/v1"
	"go-depot/pkg/limiter"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/api",      //自定义键值对名称
	FillInterval: time.Second, //间隔多久时间放 N 个令牌
	Capacity:     10,          //令牌桶的容量
	Quantum:      10,          //每次到达间隔时间后所放的具体令牌数量
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())   // gin Logger
		r.Use(gin.Recovery()) // gin recovery
	} else {
		r.Use(middleware.AccessLog()) // 自定义日志中间件
		r.Use(middleware.Recovery())  // 自定义Recovery中间件
	}
	r.Use(middleware.Translations())                                          // i18n中间件
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout)) //超时控制
	r.Use(middleware.RateLimiter(methodLimiters))                             //限流控制
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiV1 := r.Group("/api/v1")
	// 鉴权中间件
	apiV1.Use(middleware.Auth())
	// 验签中间件
	apiV1.Use(middleware.Sign())
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)
	}

	return r
}
