package router

import "github.com/gin-gonic/gin"

var origins = map[string]bool{
	"http://m.dillonliang.cn": false,
}

func Cross() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		if origins[origin] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, X-Method, X-Timestamp, X-Signature, Content-Type")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
