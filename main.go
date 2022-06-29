package main

import (
	"log"
	"net/http"
	"tour/global"
	"tour/internal/routers"
	"tour/setup"
)

/*
 init app:
      setting | db | logger
*/
func init() {
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
		global.Logger.Fatalf("main.httpServer err :%v", err)
	}
}
