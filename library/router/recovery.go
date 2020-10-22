package router

import (
	"fmt"
	"runtime/debug"

	"git.dillonliang.cn/micro-svc/pledge/library/log"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				data := []interface{}{
					"error", fmt.Sprintf("%s", err),
					"stacktrace", string(debug.Stack()),
				}
				log.Errorw("panic recovered", data...)

				c.JSON(500, gin.H{"ok": false, "code": "SERVER_ERROR"})
				c.Abort()
			}
		}()

		c.Next()
	}
}
