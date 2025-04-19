package middleware

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func MyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("key", "value")
        c.Next()
    }
}

func RestrictMethodsMiddleware(c *gin.Context) {
    if c.Request.Method != "GET" && c.Request.Method != "POST" && c.Request.Method != "OPTIONS" && c.Request.Method != "DELETE" && c.Request.Method != "PUT" {
        c.AbortWithError(405, gin.Error{Err: fmt.Errorf("Method not allowed")})
        return
    }
    c.Next()
}
