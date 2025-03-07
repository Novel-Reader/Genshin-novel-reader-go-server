package middleware

import (
	"github.com/gin-gonic/gin"
)

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在处理函数执行前添加自定义逻辑
		c.Set("key", "value")
		// 继续处理请求
		c.Next()
		// 在处理函数执行后添加自定义逻辑
	}
}
