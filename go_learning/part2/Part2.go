package main

import (
	"fmt"
	"strconv"
)

/**
  基础数据类型

   bool   true/false
   数值类型
         整数
            有符号
            int8 （-128~127）
            int16 （-2^15 ~ 2^15-1）
            int32 （-2^31 ~ 2^31-1）
            int64 （-2^64 ~ 2^64-1）

             无符号
             uint8 (0-2^8-1 )
             uint16 (0-2^16-1)
             uint32 (0-2^32-1)
             uint64 (0-2^64-1)

     	 浮点数
             float32
             float64
         复数
         字节 uint8 (
         rune int32

   字符与字符串











*/

func main() {

	//基础类型定义
	var _ int8
	var _ uint8

	var _ int //动态类型

	var _ float64 // 1.8e308
	var _ float32 // 3.4e38

	//var e float 没有这种方式

	//byte
	var _ byte //type byte = uint8

	var ch byte
	ch = 'a'
	fmt.Println(ch)
	fmt.Printf("ch = %c\n", ch)

	ch = ch + 1
	fmt.Printf("ch = %c", ch)

	//rune 标识更多的字符中文，
	var ch2 rune
	//ch = '我'
	ch2 = '我'

	//
	//constant 25105 overflows byte
	//fmt.Println(ch)
	//fmt.Printf("ch = %c\n",ch)

	fmt.Println(ch2)
	fmt.Printf("ch2 = %c", ch2)

	// 数据类型的装换

	//浮点数
	a := 5.0

	// 转化为int
	_ = int(a)

	//能将底层相同结构互换
	type IT int

	var c IT = 2
	// IT to int
	_ = int(c)
	//int to IT
	_ = IT(3)

	var d int8 = 3
	_ = uint8(d)

	//字符串转数字 工具包strconv
	i, err := strconv.Atoi("123\n")
	if err != nil {
		fmt.Println(i)
	}

	str := strconv.Itoa(123)
	fmt.Println(str)

	//字符串转化为float32/bool类型
	float, err := strconv.ParseFloat("3.1415926", 32)
	if err != nil {
		return
	}
	fmt.Println(float)

	//转化为10进制
	h, err := strconv.ParseInt("-12", 10, 64)
	if err != nil {
		return
	}
	fmt.Println(h)

	f, err := strconv.ParseInt("12", 16, 64)
	if err != nil {
		return
	}
	fmt.Println(f)

	// bool
	parseBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(parseBool)
	}
	fmt.Println(parseBool)

	parseBool1, err := strconv.ParseBool("true1")
	if err != nil {
		fmt.Println(parseBool1)
	}

	parseBool2, err := strconv.ParseBool("1")
	if err != nil {
		return
	}
	fmt.Println(parseBool2)

	//基础类型转字符串
	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool)

	formatFloat := strconv.FormatFloat(3.1425, 'E', -1, 64)
	fmt.Println(formatFloat)

	formatFloat2 := strconv.FormatFloat(3.1425, 'f', -1, 64)
	fmt.Println(formatFloat2)

	formatInt := strconv.FormatInt(12, 16)
	fmt.Println(formatInt)

	// 运算法与表达式

}
