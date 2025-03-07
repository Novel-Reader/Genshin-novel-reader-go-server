package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// 自定义中间件函数
func myMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 在处理函数执行前添加自定义逻辑
        c.Set("key", "value")
        // 继续处理请求
        c.Next()
        // 在处理函数执行后添加自定义逻辑
    }
}

func main() {
    r := gin.Default()

    // 使用全局中间件
    r.Use(myMiddleware())

    r.GET("/", func(c *gin.Context) {
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
    })

    r.Run(":8080")
}
