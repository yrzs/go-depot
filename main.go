package main

import (
	"context"
	"go-depot/global"
	"go-depot/internal/routers"
	"go-depot/setup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	/**
		优雅重启和停止:
		不关闭现有连接（正在运行中的程序）
	 	新的进程启动并替代旧进程
	 	新的进程接管新的连接
	 	连接要随时响应用户的请求。当用户仍在请求旧进程时要保持连接，新用户应请求新进程，不可以出现拒绝请求的情况
	*/
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			global.Logger.Fatalf(nil, "main.httpServer err :%v", err)
		}
	}()
	// 等待中断信号
	// 如果没有正在处理的旧请求，那么在接收到SIGINT/SIGTERM后，其会直接退出（因为不需要等待）
	quit := make(chan os.Signal)
	// 接收 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	// 最大时间控制，用于通知该服务端它有 5 秒的时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		global.Logger.Fatalf(ctx, "Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting!")
}
