package middleware

import (
	"github.com/gin-gonic/gin"
)

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("key", "value")
		c.Next()
	}
}
