package main

import (
	"container/list"
	"fmt"
	"strconv"
)

/*
   集合类型
   数组、切片、map、list

*/

func main() {

	//数组 定义 var name [count]int

	var arr [3]string // 只有三个数组的类型
	var arr2 [4]string

	//arr == arr2 //不同类型

	fmt.Printf("%T,%T \n", arr, arr2)

	arr[0] = "i"
	arr[1] = "am"
	arr[2] = "dog"

	arr2[0] = "i"

	for key, value := range arr {
		fmt.Println(key, value)
	}

	for key, value := range arr2 {
		fmt.Println(key, value)
	}

	//数组的初始化
	arr3 := [3]string{"a", "b", "c"}
	for key, value := range arr3 {
		fmt.Println(key, value)
	}

	arr4 := [3]string{1: "hehhe"}
	for key, value := range arr4 {
		fmt.Println(key, value)
	}

	arr5 := [...]string{"a", "b", "c", "dd"}
	for key, value := range arr5 {
		fmt.Println(key, value)
	}

	//
	for i := 0; i < len(arr5); i++ {
		fmt.Println(arr5[i])
	}

	//数组比较 类型一样

	//if arr2 == arr3{} //

	if arr == arr3 {
		fmt.Println("=")
	} else {
		fmt.Println("!=")
	}

	//多维数组

	var arr6 [2][3]string
	arr6[0] = [3]string{"1", "2", "3"}
	arr6[1] = [3]string{"4", "5", "6"}

	fmt.Println(len(arr6))

	for i := 0; i < len(arr6); i++ {
		for j := 0; j < len(arr6[i]); j++ {
			fmt.Print(arr6[i][j] + "")
		}
		fmt.Println("")
	}

	for _, row := range arr6 {
		for _, v1 := range row {
			fmt.Print(v1)
		}
		fmt.Println(" ")
	}

	// 切片（slice） --类似动态数组
	// 定义 var xxx []string
	var s1 []string

	fmt.Printf("%T \n", s1)

	s1 = append(s1, "i")
	s1 = append(s1, "am")
	s1 = append(s1, "golang")

	fmt.Println(s1[1])

	//切片的初始化
	//1 数组创建
	//2 使用type{}
	//3 make

	//
	arr7 := [3]string{"a", "b", "c"}

	s2 := arr7[0:len(arr7)] //左闭右开
	fmt.Printf("%T,%T \n  ", arr7, s2)

	//
	s3 := []string{"a", "b", "c"}
	fmt.Printf("%T,%T \n  ", arr7, s3)

	//
	s4 := make([]string, 3)
	s4[0] = "ee"
	s4[1] = "e1"
	fmt.Printf("%T,%T \n  ", arr7, s4)

	//
	//var s5 []string //没有初始化只能使用append
	//s5[0] = "i"

	// 访问切片的数据
	fmt.Println(s2[0])
	fmt.Println(s2[1:2])
	fmt.Println(s2[1:])
	fmt.Println(s2[:1])
	fmt.Println(s2[:])

	//数据添加
	s6 := []string{"go", "java"}
	s6 = append(s6, "rust")
	fmt.Println(s6)

	//合并
	s6 = append(s6, s2...)
	fmt.Println(s6)

	s6 = append(s6, s6[0:1]...)
	fmt.Println(s6)

	//删除
	s6 = append(s6[:2], s6[3:]...)
	fmt.Println(s6)

	//复制
	s7 := s6
	s8 := s6[:]
	fmt.Println(s7, s8)

	var s9 []string
	copy(s9, s6)
	fmt.Println(s9) //注意这个地方、对比下面的方法

	s10 := make([]string, len(s6))
	copy(s10, s6)
	fmt.Println(s10)

	//对比
	s6[0] = "update"
	fmt.Println(s8, s10) //赋值相当于浅拷贝/copy相当于深拷贝

	//分析切片原理

	// 值传递与引用传递

	s11 := []string{"a", "b", "c"}
	fmt.Println(s11)
	printSlice(s11)
	fmt.Println(s11)

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s12 := data[0:5]
	s13 := data[3:]
	s13[0] = 100
	fmt.Println(s12, s13, data) //值修改
	//[1 2 3 100 5] [100 5 6 7 8 9 10] [1 2 3 100 5 6 7 8 9 10]

	data2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s14 := data2[0:5]
	s15 := data2[3:]

	s15 = append(s15, 433) //s15 扩容底层数组已经修改
	s15[0] = 100
	fmt.Println(s14, s15, data2) //
	//[1 2 3 4 5] [100 5 6 7 8 9 10 433] [1 2 3 4 5 6 7 8 9 10]

	// Map 集合 k,v 的无序集合
	// 定义 var <xx> map[type]type,
	// 必须初始化才能放值
	// 不是线程安全的

	var iMap map[string]string
	fmt.Println(iMap)
	fmt.Println(iMap["id"])
	//iMap["id"] = "id"

	//初始化
	var m1 map[string]string = map[string]string{
		"id": "m1",
	}

	var m2 = map[string]string{
		"id": "m2",
	}

	m3 := map[string]string{
		"id":   "m3",
		"name": "",
	}

	m4 := map[string]string{}

	m5 := make(map[string]string, 8)

	m5["id"] = "m5"

	fmt.Println(m1, m2, m3, m4, m5)
	fmt.Println(m5["id"])

	//遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	for key := range m1 {
		fmt.Println(key)
	}

	//contains
	fmt.Println(m3["name"])
	fmt.Println(m4["name"])

	n1, ok := m3["name"]
	if ok {
		fmt.Println("in", n1)
	}

	n2, ok1 := m4["name"]
	if !ok1 {
		fmt.Println("Not in", n2)
	}

	//改写
	if n3, ok3 := m3["name"]; ok3 {
		fmt.Println("in", n3)
	}

	//删除
	m5["delete"] = "delete"
	fmt.Println(m5)

	delete(m5, "delete")
	fmt.Println(m5)

	// list表
	//container/list

	var myList list.List

	myList.PushBack("id")
	myList.PushBack("name")

	fmt.Println(myList)

	//正序
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	for j := myList.Back(); j != nil; j = j.Prev() {
		fmt.Println(j.Value)
	}

	//remove
	myList.Remove(myList.Back())

	for j := myList.Back(); j != nil; j = j.Prev() {
		fmt.Println(j.Value)
	}

}

func printSlice(data []string) {
	data[0] = "java"

	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}

}
