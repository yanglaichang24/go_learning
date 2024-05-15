package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	r.POST("/test1", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/test")
	})

	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test3"
		r.HandleContext(c)
	})
	r.GET("/test3", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	r.Run(":8080")
}
