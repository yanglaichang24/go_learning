package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**/
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")

		names := c.PostFormMap("names")

		fmt.Printf("-------ids: %v; names: %v---------------", ids, names)

		c.JSON(200, gin.H{
			"code": "OK",
		})

	})
	router.Run(":8080")
}
