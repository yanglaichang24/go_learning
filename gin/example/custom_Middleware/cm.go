package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

/*
 */
func filter() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("----filter starting--")
		c.Next()
		log.Print("----filter end --")
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("Logger--starting--")

		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request
		log.Print("----before request--")
		c.Next()
		log.Print("----after request--")

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
		log.Print("Logger---end--")
	}
}

func main() {
	r := gin.New()

	//类似filterChain
	r.Use(filter())
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println("----" + example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
