package setup

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"time"
	"tour/global"
	"tour/internal/model"
	"tour/pkg/logger"
	"tour/pkg/setting"
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
