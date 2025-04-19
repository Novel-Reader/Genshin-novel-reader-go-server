package index

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "text/template"
)

func IndexRoutes(r *gin.RouterGroup) {
    r.GET("/", func(c *gin.Context) {
        t, err := template.New("welcome").ParseFiles("templates/welcome.html")
        if err != nil {
            c.String(http.StatusInternalServerError, err.Error())
            return
        }
        data := struct {
            Title string
            Body  string
        }{
            Title: "Genshin novel reader go server",
            Body:  "Welcome to use Genshin novel reader go server",
        }
        t.ExecuteTemplate(c.Writer, "welcome.html", data)
    })
}
