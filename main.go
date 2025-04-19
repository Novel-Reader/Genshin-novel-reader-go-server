package main

import (
    "log"
    "os"
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
    // 打开日志文件
    // os.OpenFile() 函数有三个参数：
    // 1、path 文件路径
    // 2、flag 文件模式，表示以只写模式打开文件，如果文件不存在则创建一个新文件，所有写入操作都将追加到文件的末尾。
    // 3、mode 文件权限，可以设置为 0666，表示所有人都可以读写
    logFile, err := os.OpenFile("novel-reader-go-server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        // log.Fatal 通常用于处理那些无法恢复的错误，例如文件无法打开、网络连接失败等情况。
        // 记录一个错误消息到标准错误输出（stderr）；终止程序的执行，返回一个非零的退出状态码。
        log.Fatal(err)
    }
    // 设置日志输出的格式：日期, 时间, 文件名和行号（短格式）
    // 2025/01/01 08:00:00 example.go:12: Hello, World!
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetOutput(color.Output)

    logFile, err := os.OpenFile("novel-reader-go-server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetOutput(color.Output)

    // 设置框架的日志输出文件
    gin.DefaultWriter = logFile
    gin.DefaultErrorWriter = logFile

    // 加载配置文件
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
