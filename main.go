package main

import (
    "github.com/gin-gonic/gin"
    "myproject/middleware"
    "myproject/routes"
    "myproject/config"
    "log"
    "os"
    "github.com/fatih/color"
    "time"
    "github.com/gin-contrib/cors"
)

func main() {

    logFile, err := os.OpenFile("novel-reader-go-server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetOutput(color.Output)

    gin.DefaultWriter = logFile
    gin.DefaultErrorWriter = logFile

    config, err := config.LoadConfig("config/config.json")
    if err != nil {
        log.Fatal(err)
    }

    r := gin.Default()
    r.SetTrustedProxies([]string{"127.0.0.1"})

    // middleware
    r.Use(middleware.RestrictMethodsMiddleware)

    // CORS
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowAllOrigins = true
    corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
    corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
    corsConfig.AllowCredentials = true
    corsConfig.MaxAge = 12 * time.Hour
    r.Use(cors.New(corsConfig))

    r.Use(middleware.MyMiddleware())

    // routes
    routes.Routes(&r.RouterGroup)

    r.Run(":" + config.Port)
}

// run `go run main.go` to start server
