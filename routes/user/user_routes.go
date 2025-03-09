package user

import (
	"github.com/gin-gonic/gin"
	"github.com/fatih/color"
	"net/http"
	"log"
	"myproject/db"
)

func UserRoutes(r *gin.RouterGroup) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(color.Output)

	r.GET("/api/user", getUser)
	r.POST("/api/user", createUser)
	r.PUT("/api/user", updateUser)
	r.DELETE("/api/user", deleteUser)
}

func getUser(c *gin.Context) {
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
}

func createUser(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	email := c.PostForm("email")
	name := c.PostForm("name")
	password := c.PostForm("password")

	if email == "" || name == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if email == "" || name == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email, username or password is not correct"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is too short"})
		return
	}
	results, err := db.QueryDB("SELECT * FROM user WHERE email=?", email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if len(results) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	// If email not found, insert new email into db
	sql := "INSERT INTO user (name, email, password) VALUES (?, ?, ?)"
	_, err = db.QueryDB(sql, name, email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert new user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User inserted successfully"})
}

func updateUser(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"message": "Not allowed",
	})
}

func deleteUser(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}
	_, err := db.QueryDB("DELETE FROM user WHERE email = ?", email)
	if err != nil {
			log.Println(err)
			return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete user successfully",
	})
}
