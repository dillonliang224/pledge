package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.dillonliang.cn/micro-svc/pledge/library/log"
	"git.dillonliang.cn/micro-svc/pledge/src/base/book/api"
	"git.dillonliang.cn/micro-svc/pledge/src/base/book/config"
	"git.dillonliang.cn/micro-svc/pledge/src/base/book/service"
	"git.dillonliang.cn/micro-svc/pledge/src/base/book/transport/grpc"
)

func main() {
	cfg := config.Load(api.App)

	logger := log.SetLogger(api.App, cfg.Log.Level)

	svc := service.New(cfg, logger)
	_, err := grpc.Start(cfg, svc)
	if err != nil {
		panic(fmt.Sprintf("start base book failed: %v", err))
	}

	log.Info("base book service start listen on ", cfg.Port.GRPC)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Warnw("base book service exit")
			svc.Close()

			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
