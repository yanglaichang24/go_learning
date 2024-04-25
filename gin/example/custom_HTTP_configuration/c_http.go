package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
 */
func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}

}

/*

also user next

func main() {
	router := gin.Default()
	http.ListenAndServe(":8080", router)
}


*/
