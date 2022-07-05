package setup

import (
	"go-depot/global"
	"go-depot/internal/model"
	"go-depot/pkg/logger"
	"go-depot/pkg/setting"
	"go-depot/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"time"
)

/**
mapping setting
*/
func Setting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	err = s.ReadSection("ApiClient", &global.ApiClientSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Wechat", &global.WechatSetting)
	if err != nil {
		return err
	}
	return nil
}

/**
setup db database
*/
func DB() error {
	var err error
	/**
	if use：< global.DB, err := model.NewDB(global.DatabaseSetting) > is terrible
	因为 := 会重新声明并创建了左侧的新局部变量，因此在其它包中调用 global.DB 变量时，它仍然是 nil，仍然是达不到可用标准
	因为根本就没有赋值到真正需要赋值的包全局变量 global.DB 上
	*/
	global.DB, err = model.NewDB(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

/**
setup logger
*/
func Logger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

/**
setup tracer(jaeger)
*/
func Tracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		global.AppSetting.OpenTracing.ServiceName,
		global.AppSetting.OpenTracing.AgentHost+":"+global.AppSetting.OpenTracing.AgentPort,
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
