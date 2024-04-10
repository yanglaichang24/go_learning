package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("test.txt") // 打开文件
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // 确保文件在函数结束时关闭

	reader := bufio.NewReader(file) // 创建新的读取器

	for {
		line, err := reader.ReadString('\n') // 读取一行数据

		fmt.Printf("Line: %s", line) // 打印读取到的数据

		if err == io.EOF { // 如果遇到EOF，则退出循环
			break
		}

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
