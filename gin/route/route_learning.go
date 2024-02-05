package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	//路由区分
	group := route.Group("/v1")
	{
		group.GET("/ping", pong)
	}

	group2 := route.Group("/v2")
	{
		group2.GET("/ping", pong2)
	}

	//获取path中的数据
	group3 := route.Group("/v3")
	{
		group3.GET("/ping/:id", pong3)
		group3.GET("/ping/:id/get", pong3)

		//http://localhost:8081/v3/ping2/id/get/name/java/go
		//异常
		group3.GET("/ping/:id/get/:name", pong4)

		//http://localhost:8081/v3/ping2/id/get/name/java/go
		// name= "/get/name/java/go"
		group3.GET("/ping2/:id/*name", pong4)
	}

	//path 约束
	route.GET("/user/:id/:name", func(c *gin.Context) {
		var person Person
		//这里好几个方法，注意别整错了
		if err := c.ShouldBindUri(&person); err != nil {
			c.Status(404)
			println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":   person.Id,
			"name": person.Name,
		})
	})

	route.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type Person struct {
	//Id   string `uri:"id" binding:"required"`
	Id   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
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

func pong3(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"id":      id,
	})
}

func pong4(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"id":      id,
		"name":    name,
	})
}
