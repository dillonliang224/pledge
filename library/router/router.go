package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"git.dillonliang.cn/micro-svc/pledge/library/config"
)

const (
	UserKey = "userId"
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

func (r *Router) AuthUser(c *gin.Context) {
	userId := r.GetUserId(c)
	if userId == "TODO" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"ecode": 0,
		})
		return
	}

	c.Set(UserKey, userId)
}

func (r *Router) GetUserId(c *gin.Context) (userId string) {
	token := c.Query("token")
	return token
}
