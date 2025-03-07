package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fatih/color"
	"log"
	"myproject/db"
)

func IndexRoute(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html><body>Welcome to use Genshin novel reader go server</body></html>"))
}

func PingRoute(c *gin.Context) {
	if c.Request.Method == "GET" {
			c.JSON(http.StatusOK, gin.H{
					"message": "pong",
			})
	} else if c.Request.Method == "POST" {
			c.JSON(http.StatusOK, gin.H{
					"message": "pong",
			})
	} else {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
					"message": "Not allowed",
			})
	}
}

func ApiUserRoute(c *gin.Context) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(color.Output)
	if c.Request.Method == "GET" {
		email := c.Query("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
			return
		}
		results, err := db.QueryDB("SELECT * FROM user WHERE email = ? LIMIT 1", email)
		if err != nil {
				log.Println(err)
				return
		}
		for _, row := range results {
				log.Println(email, row["id"], row["name"], row["email"])
		}
		if len(results) > 0 {
			// 数据库中是 ASCII 编码，需要转换成 UTF-8 编码
			nameBytes := results[0]["name"].([]byte)
			nameString := string(nameBytes)
			emailBytes := results[0]["email"].([]byte)
			emailString := string(emailBytes)
			c.JSON(http.StatusOK, gin.H{
				"id":    results[0]["id"],
				"name":  nameString,
				"email": emailString,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		}
	} else if c.Request.Method == "POST" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Not support now",
		})
	} else if c.Request.Method == "DELETE" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Not support now",
		})
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Not allowed",
		})
	}
}
