package main

import "github.com/gin-gonic/gin"

/*
  Bind html checkboxes
*/

type myForm struct {
	//切片
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	_ = c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

func main() {
	r := gin.Default()
	r.GET("/checkbox", formHandler)
	_ = r.Run()

}
