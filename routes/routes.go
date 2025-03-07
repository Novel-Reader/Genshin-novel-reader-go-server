package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexRoute(c *gin.Context) {
	// 从上下文中获取中间件设置的值
	value, exists := c.Get("key")
	if exists {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World! " + value.(string),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	}
}
