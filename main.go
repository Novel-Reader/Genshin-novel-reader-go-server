package main

import (
    "github.com/gin-gonic/gin"
		"myproject/middleware"
		"myproject/routes"
		"myproject/config"
		"log"
)

func main() {

		config, err := config.LoadConfig("config/config.json")
		if err != nil {
			log.Fatal(err)
		}

    r := gin.Default()
		r.Use(middleware.MyMiddleware())
		r.Any("/", routes.IndexRoute)
		r.Any("/ping", routes.PingRoute)
		r.Any("/api/user", routes.ApiUserRoute)
		r.Run(":" + config.Port)
}

// run `go run main.go` to start server
