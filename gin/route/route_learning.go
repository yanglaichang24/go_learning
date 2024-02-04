package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	group := route.Group("/v1")
	{
		group.GET("/ping", pong)
	}

	group2 := route.Group("/v2")
	{
		group2.GET("/ping", pong2)
	}

	route.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func pong2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong2",
	})
}
