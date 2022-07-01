package global

import (
	"go-depot/pkg/logger"
	"go-depot/pkg/setting"
)

var (
	ServerSetting    *setting.ServerSettingS
	AppSetting       *setting.AppSettingS
	DatabaseSetting  *setting.DatabaseSettingS
	ApiClientSetting *setting.ApiClientSettingS
	WechatSetting    *setting.WechatSettingS
	Logger           *logger.Logger
)
