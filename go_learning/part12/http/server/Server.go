package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler 响应HTTP GET请求，返回"Hello, World!"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// 设置路由规则
	http.HandleFunc("/", helloHandler)

	// 启动HTTP服务器，监听8080端口
	fmt.Println("Starting HTTP server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
