package service

import (
	"go-depot/global"
	"go-depot/pkg/convert"
	"go-depot/pkg/errcode"
	"strings"
	"time"
)

type ApiAccessTokenRequest struct {
	AccessToken string `form:"access_token" binding:"required"`
}

/**
validate api x-api-key
*/
func (svc Service) CheckAuth(param *ApiAccessTokenRequest) *errcode.Error {
	if !global.ApiClientSetting.AccessTokenValidity {
		return errcode.Success
	}
	acInfo, err := svc.dao.GetApiAccessTokenInfoByAccessToken(param.AccessToken)
	if err != nil || acInfo == nil {
		return errcode.UnauthorizedAuthNotExist
	}
	//截取过期时间
	accessTokenArr := strings.Split(acInfo.AccessToken, "_")
	expire := accessTokenArr[len(accessTokenArr)-1]
	if convert.StrTo(expire).MustInt() > (int(time.Now().Unix()) + global.ApiClientSetting.AccessTokenExpire) {
		return errcode.UnauthorizedTokenTimeout
	}
	return errcode.Success
}
