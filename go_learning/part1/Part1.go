package main

import "fmt"

/*
* 定义变量（全局变量）
 */

// 定义全局变量
var a int
var b = "b"
var flag bool

//c:= "c" // 不能用

var (
	c string
	d = 1
)

func main() {

	//定义局部变量
	//局部变量必须使用
	var name int
	fmt.Println(name)

	var id = 1
	fmt.Println(id)

	var key = "star"
	fmt.Println(key)

	short := "ll"
	fmt.Println(short)

	// 多变量定义
	var x, y, z string
	fmt.Println(x, y, z)

	var o, p, q = "o", 2, "q"
	fmt.Println(o, p, q)

	//定义常量
	const PI = 3.1415926
	const PI_2 float64 = 3.1415926

	const (
		P_1 = 1
		P_2 = 2
	)

	const (
		P_3 int = 1
		P_4     = 2
		P_5
		P_6 = "7"
		P_7
	)

	// NOTE: 注意P_5/P_6的值
	fmt.Println(P_3, P_4, P_5, P_6, P_7)

	// 特殊常量iota
	const (
		E1 = iota
		//E1 = iota + 1
		E2
		E3

		E4 = "err"
		E5
		E6 = iota
		E7
		E8 = iota + 1
		E9
	)

	// 重新开始了
	const (
		E10 = iota
	)

	fmt.Println("-------------------")
	fmt.Println(E1, E2, E3, E4, E5, E6, E7, E8, E9, E10)

	// 匿名变量可以不用使用，对于一些函数的返回值不使用可以使用匿名
	var _ int

	// 变量的作用域
	{
		localvar := "lo"
		fmt.Println(localvar)
	}
	//fmt.Println(localvar)

	var yang string

	// 定义在外围或者全局变量
	var localvar string
	if yang == "yang" {
		localvar := "l1"
		fmt.Println(localvar)
	}
	fmt.Println(localvar)

}
