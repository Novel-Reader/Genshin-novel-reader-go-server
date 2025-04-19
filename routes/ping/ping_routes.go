package ping

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func PingRoutes(r *gin.RouterGroup) {
    r.GET("/ping", pingGet)
    r.POST("/ping", pingPost)
    r.PUT("/ping", pingOther)
    r.DELETE("/ping", pingOther)
}

func pingGet(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}

func pingPost(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}

func pingOther(c *gin.Context) {
    c.JSON(http.StatusMethodNotAllowed, gin.H{
        "message": "Not allowed",
    })
}
