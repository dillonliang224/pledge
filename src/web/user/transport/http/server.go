package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"git.dillonliang.cn/micro-svc/pledge/library/router"
	"git.dillonliang.cn/micro-svc/pledge/src/web/user/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/web/user/service"
)

func Start(cfg *conf.Config, service *service.Service) {
	r := router.New(cfg.Common)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"msg": "hello",
		})
	})

	go func() {
		if err := r.Run(cfg.Port.HTTP); err != nil {
			panic(err)
		}
	}()
}

type server struct {
	as *service.Service
}
