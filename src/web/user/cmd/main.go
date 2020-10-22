package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.dillonliang.cn/micro-svc/pledge/library/log"
	"git.dillonliang.cn/micro-svc/pledge/src/web/user/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/web/user/service"
	"git.dillonliang.cn/micro-svc/pledge/src/web/user/transport/http"
)

const App = "web-user"

func main() {
	cfg := conf.Load(App)
	logger := log.SetLogger(App, cfg.Log.Level)

	svc := service.New(cfg, logger)
	http.Start(cfg, svc)

	log.Info("web user service start listen on ", cfg.Port.HTTP)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Warnw("web user service exit")
			svc.Close()

			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
