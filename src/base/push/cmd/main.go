package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.dillonliang.cn/micro-svc/pledge/library/log"
	"git.dillonliang.cn/micro-svc/pledge/src/base/push/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/base/push/service"
	"git.dillonliang.cn/micro-svc/pledge/src/base/push/transport/grpc"
)

const App = "base-push"

func main() {
	cfg := conf.Load(App)
	logger := log.SetLogger(App, cfg.Log.Level)

	svc := service.New(cfg, logger)
	_, err := grpc.Start(cfg, svc)
	if err != nil {
		panic(fmt.Sprintf("start base push failed: %v", err))
	}

	log.Info("base push service start listen on ", cfg.Port.GRPC)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Warnw("base push service exit")
			svc.Close()

			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
