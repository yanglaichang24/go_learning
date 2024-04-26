package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 */
func main() {
	r := gin.Default()
	r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("templates/delims/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "delims.tmpl", gin.H{
			"title": "Main website########",
		})
	})
	r.Run()

}
