package routers

import (
	"github.com/gin-gonic/gin"
	"go-depot/internal/middleware"
	v1 "go-depot/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	// gin Logger
	r.Use(gin.Logger())
	// 自定义日志中间件
	r.Use(middleware.AccessLog())
	// gin recovery
	r.Use(gin.Recovery())
	// 自定义Recovery中间件
	r.Use(middleware.Recovery())
	// i18n中间件
	r.Use(middleware.Translations())

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
