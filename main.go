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
		r.GET("/", routes.IndexRoute)
		r.Run(":" + config.Port)
}

// run `go run main.go` to start server
