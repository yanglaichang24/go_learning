package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 字符串操作
*/

func main() {

	// 长度计算
	name := "我是长度计算"
	bytes := []rune(name)
	fmt.Println(len(bytes))

	//转义字符
	a := "我是\"大猩猩\""
	fmt.Println(a)

	b := `我是"大猩猩"`
	fmt.Println(b)

	//格式化输出
	username := "yanglaichang"
	age := 19
	addr := "beijignshi"

	fmt.Println("姓名=" + username + ",年龄：" + strconv.Itoa(age) + ",地址：" + addr)
	fmt.Printf("姓名=：%s,年龄：%d,地址：%s\t\n", username, age, addr)
	sprintf := fmt.Sprintf("姓名：%s,年龄：%d,地址：%s", username, age, addr)
	fmt.Println(sprintf)

	//stringBuilder
	var builder strings.Builder
	builder.WriteString("-------姓名：")
	builder.WriteString(username)
	builder.WriteString("年龄：")
	builder.WriteString(strconv.Itoa(age))

	build := builder.String()
	fmt.Println(build)

}
