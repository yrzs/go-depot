package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-depot/global"
	"go-depot/pkg/app"
	"go-depot/pkg/convert"
	"go-depot/pkg/encrypt"
	"go-depot/pkg/errcode"
	"go-depot/pkg/utils"
	"sort"
	"time"
)

/**
接口验签中间件
*/
func Sign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 判断是否开启接口验签
		if global.ApiClientSetting.HttpSignValidity {
			var ecode = errcode.Success
			if ctx.Request.Form == nil {
				ctx.Request.ParseMultipartForm(32 << 20)
			}
			queryMap := make(map[string]string)
			var sign, signStr, calcSign string
			var keys []string
			for k, v := range ctx.Request.Form {
				fmt.Println(k, v[0])
				if k == "__sign" {
					sign = v[0]
				} else {
					queryMap[k] = v[0]
					keys = append(keys, k)
				}
			}
			if time.Now().Unix()-convert.StrTo(queryMap["__time"]).MustInt64() > global.ApiClientSetting.HttpSignExpire {
				ecode = errcode.SignTimeOut
			}
			if !utils.IsEmpty(keys) {
				//排序 ascii
				sort.Strings(keys)
				for _, key := range keys {
					signStr += key + "=" + queryMap[key] + "&"
				}
				// 去掉最后一个&
				signStr = signStr[:len(signStr)-1]
				// 排好序的参数加上secret,进行md5
				signStr += global.ApiClientSetting.HttpSignAccount.Secret
				calcSign = encrypt.MD5(signStr)
				if calcSign != sign {
					ecode = errcode.SignError
				}
			} else {
				ecode = errcode.SignError
			}
			if ecode != errcode.Success {
				response := app.NewResponse(ctx)
				response.ToErrorResponse(ecode)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
