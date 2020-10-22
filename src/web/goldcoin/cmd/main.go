package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.dillonliang.cn/micro-svc/pledge/library/log"
	"git.dillonliang.cn/micro-svc/pledge/library/router/debug"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/service"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/transport/http"
)

const (
	app = "web-goldcoin"
)

func main() {
	cfg := conf.Load(app)
	log.SetLogger(app, cfg.Log.Level)

	svc := service.New(cfg)
	http.Start(cfg, svc)
	debug.Start(cfg.Common, svc)

	log.Info("web goldcoin service started listen on ", cfg.Port.HTTP)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Warnw("web goldcoin service exit")
			svc.Close()

			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
