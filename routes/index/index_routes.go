package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexRoutes(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html><body>Welcome to use Genshin novel reader go server</body></html>"))
	})
}
