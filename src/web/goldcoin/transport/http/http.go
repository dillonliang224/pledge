package http

import (
	"git.dillonliang.cn/micro-svc/pledge/library/router"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/service"
)

func Start(c *conf.Config, svc *service.Service) {
	r := router.New(c.Common)

	// start http server
	go func() {
		if err := r.Run(c.Common.Port.HTTP); err != nil {
			panic(err)
		}
	}()
}
