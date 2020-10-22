package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"git.dillonliang.cn/micro-svc/pledge/library/config"
)

type Router struct {
	*gin.Engine
}

func New(c config.Common) *Router {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(Cross(), Recovery())

	engine.Use(func(c *gin.Context) {
		c.Next()

		if err := c.Errors.Last(); err != nil {
			status := http.StatusOK

			c.JSON(status, gin.H{
				"ok":      false,
				"message": err.Error(),
			})
		}
	})

	return &Router{Engine: engine}
}
