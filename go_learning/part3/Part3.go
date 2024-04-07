package main

import "fmt"

/*运算符

	数值运算符
	关系运算符
    逻辑运算符
	位运算符
	赋值运算符
*/

func main() {

	//数值运算符 + - * / ++ --
	var a, b = 1, 2
	fmt.Println(a + b)

	var c, d = "c", "d"
	fmt.Println(c + d)

	fmt.Println(3 % 2)
	//关系运算符

	//逻辑运算符
	var e, f = true, false
	fmt.Println(e && f)

	//位运算符
	fmt.Println(1 & 1)
	fmt.Println(1 | 1)

	fmt.Println(1 ^ 1)
	fmt.Println(0 ^ 0)

	//赋值运算符
	h := 10
	h += 5
	fmt.Println(h)

	h >>= 3

	fmt.Println(h)

}
