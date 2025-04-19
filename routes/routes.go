package routes

import (
    "github.com/gin-gonic/gin"
    "myproject/routes/index"
    "myproject/routes/ping"
    "myproject/routes/user"
)

func Routes(r *gin.RouterGroup) {
    index.IndexRoutes(r)
    ping.PingRoutes(r)
    user.UserRoutes(r)
}
