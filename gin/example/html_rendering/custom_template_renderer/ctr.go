package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

/*
 */
func main() {

	gin.ForceConsoleColor()
	router := gin.Default()

	html := template.Must(template.ParseFiles("templates/index.tmpl", "templates/posts/index.tmpl"))
	router.SetHTMLTemplate(html)

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run()

}
