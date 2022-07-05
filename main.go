package main

import (
	"go-depot/global"
	"go-depot/internal/routers"
	"go-depot/setup"
	"log"
	"net/http"
)

/*
 init app:
      setting | db | logger
*/
func init() {
	_ = setup.Flag() // 配置信息出错了也不要紧
	err := setup.Setting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setup.DB()
	if err != nil {
		log.Fatalf("init.setupSettingDB err: %v", err)
	}
	err = setup.Logger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setup.Tracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		global.Logger.Fatalf(nil, "main.httpServer err :%v", err)
	}
}
