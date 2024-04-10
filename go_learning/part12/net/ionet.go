package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://example.com") // 发送HTTP GET请求
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() // 确保响应体在函数结束时关闭

	body, err := ioutil.ReadAll(resp.Body) // 读取响应体
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body)) // 打印响应内容
}
