package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// GET
	r.GET("/form", getDemo)

	r.POST("/form/post", postDemo)

	r.Run(":8083") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func postDemo(c *gin.Context) {
	id := c.PostForm("id")
	name := c.DefaultPostForm("name", "default")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"id":   id,
	})
}

func getDemo(c *gin.Context) {

	//
	name := c.DefaultQuery("name", "yang")
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"id":   id,
	})
}
