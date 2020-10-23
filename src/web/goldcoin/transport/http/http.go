package http

import (
	"git.dillonliang.cn/micro-svc/pledge/library/router"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/service"
)

var (
	svc *service.Service
)

func Start(c *conf.Config, s *service.Service) {
	svc = s

	r := router.New(c.Common)

	treasure := r.Group("/treasure", r.AuthUser)
	{
		treasure.GET("/info", getTreasureInfo)
		treasure.GET("/my-records", getTreasureUserRecords)
		treasure.POST("/codes", postTreasureCodes)
		treasure.GET("/participation", getTreasureParticipation)
		treasure.GET("/past-winners", getTreasureWinners)
		treasure.GET("/my-participation", getTreasureMyParticipation)
		treasure.GET("/detail-info", getTreasureDetailInfo)
		treasure.GET("/records", getTreasureRecords)
		treasure.GET("/pop", getTreasurePop)
		treasure.GET("/client-notify", getTreasureClientNotification)
		treasure.GET("/statistics", getTreasureStatistics)
	}

	// start http server
	go func() {
		if err := r.Run(c.Common.Port.HTTP); err != nil {
			panic(err)
		}
	}()
}
