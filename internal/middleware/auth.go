package middleware

import (
	"github.com/gin-gonic/gin"
	"go-depot/global"
	"go-depot/internal/service"
	"go-depot/pkg/app"
	"go-depot/pkg/errcode"
)

/**
auth middleware
*/
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			accessToken string
			errorCode   = errcode.Success
		)
		accessToken = ctx.GetHeader(global.ApiClientSetting.AccessTokenIdentity)
		if accessToken == "" {
			return
			//errorCode = errcode.InvalidParams
		}
		svc := service.New(ctx.Request.Context())
		var param = &service.ApiAccessTokenRequest{
			AccessToken: accessToken,
		}
		err := svc.CheckAuth(param)
		if err != nil {
			errorCode = err
		}
		if errorCode != errcode.Success {
			response := app.NewResponse(ctx)
			response.ToErrorResponse(errorCode)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
