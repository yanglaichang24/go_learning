package main

import (
	"fmt"
	"html/template"
)

/**/
func main() {

	if _, e := template.ParseFiles("learning-go.md"); e != nil {
		fmt.Println("====================" + e.Error())
	}

}
