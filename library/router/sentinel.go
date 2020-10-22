package router

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/gin-gonic/gin"
)

func Sentinel(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := sentinel.InitDefault()
		if err != nil {

		}
	}
}
