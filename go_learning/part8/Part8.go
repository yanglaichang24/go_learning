package main

import (
	"fmt"
	"strconv"
)

/*
  结构体
   type关键字

    定义结构体 struct
    定义接口
    定义类型别名
    类型定义

*/

type MyType int

func (mi MyType) string() string {
	return strconv.Itoa(int(mi))
}

func main() {

	/*
	 ======================================================
	                   类型别名
	 ======================================================
	*/

	//定义类型别名
	type MyInt = int
	var i MyInt = 12
	var j int = 8

	fmt.Println(i + j)
	fmt.Printf("%T \n ", i)

	/*
	  ================================================
	                   自定义类型
	  =================================================
	*/

	type MyInt2 int //自定义类型//希望int 类型，可以自定义方法,类似包装类，
	var x MyInt2 = 12

	fmt.Println(int(x) + j)
	fmt.Printf("%T \n", x)

	var m3 MyType = 13
	fmt.Println(m3.string())

	/*
	  ====================================================================
	                                 结构体
	  ====================================================================
	*/
	//结构体
	//定义一个二维切片
	var persons [][]string
	persons = append(persons, []string{"age", "addr", "name"})

	p1 := Person{}

	p2 := Person{
		name: "lll",
		age:  12,
		addr: "china",
	}

	p3 := Person{"lll", 12, "china"}

	var p4 []Person
	p4 = append(p4, p1)
	p4 = append(p4, p2)

	p4 = append(p4, Person{
		name: "xxx",
		age:  0,
		addr: "ttttt",
	})

	p5 := []Person{
		{"xxx", 0, "xx"},
		{"ddd", 0, "ddd"},
	}

	var p6 Person

	p6.name = "xxx"
	p6.age = 12
	fmt.Println(p6.name)
	fmt.Println(p1, p2, p3, p4, p5)

	//匿名结构体(一次性结构体)
	address := struct {
		provice string
		city    string
	}{
		"xxxxx",
		"yyyy",
	}
	fmt.Println(address.city)

	// 嵌套
	/*  s1 := Student{
		 p:     Person{"",0,""},
		 score: 0,
	 }

	s1.p.addr ="china"
	fmt.Println(s1,s1.p.name)*/

	s2 := Student{}
	s2.name = "xxxxx"
	fmt.Println(s2)

	s2.toString()

	s2 = Student{
		Person{name: "xxx", age: 0, addr: "xdfdf"},
		0, "name",
	}

	fmt.Println(s2.name)

	s2.name = "setname"
	fmt.Println(s2.name)
	fmt.Println(s2)

	// 结构体绑定方法
	// func(s structType) funcName(para1 pType....) (returnType)
	p3.toString()
	fmt.Println(p3.name)

	p3.print()
	fmt.Println(p3.name)

	s2.toString()

	/*
	 ========================================================================
	                             指针
	 =======================================================================
	*/

	//指针 指针需要初始化
	//函数中修改
	//指针定义 var xxx *type

	p7 := Person{
		name: "xxxx",
		age:  0,
		addr: "xxx",
	}

	p8 := Person{
		name: "xxxx",
		age:  0,
		addr: "xxx",
	}

	updateName(&p7)
	updateName2(p8)

	fmt.Println(p7.name, p8.name)

	//定义
	var p9 *Person
	p9 = &p8 //将p8的地址放进p9中
	fmt.Printf("%p,%p", p9, &p8)

	var p10 *Person = &Person{"", 0, "xxx"}

	//访问方式
	(*p10).name = "xxx"
	p10.name = "yyy"
	// go中不支持指针计算

	fmt.Println(p10)

	/*
	           P8                                   p9
	        --------------------                -------------------
	        -  {name,age,addr}  - <----------- -   0xc00006e570   -
	        --------------------  <------       -------------------
	                                    -
	                                    -
	          P8'                       -            p9'
		     ----------------------     -       -------------------
		     -  {{name,age,addr}} -     ---------  0xc00006e570   -
		     ----------------------             -------------------

	      根据函数值传递也就是copy数据
	      普通变量的复制是拷贝一份全部的数据类似p8'
	      指针变量的复制是拷贝一份地址p9'
	      所以通过指针变量可以修改原始数据

	*/
	fmt.Println("-------------------")
	//指针赋值
	y := 10
	var z *int = &y

	fmt.Println(10 + y)
	fmt.Println(10 + *z)

	fmt.Println(&z, z)

	//指针需要初始化
	var p11 *Person
	//fmt.Println(p11.name)
	// panic: runtime error: invalid memory address or nil pointer dereference

	//第一种初始化
	p11 = &Person{}
	fmt.Println(p11.name)

	//第二种
	var p12 Person
	p11 = &p12
	fmt.Println(p11.name)

	// 第三种(推荐)
	p11 = new(Person)
	fmt.Println(p11.name)

	//交换两个值
	m := 10
	n := 100
	swap(&m, &n)
	fmt.Println(m, n)

	/*
	   ===============================================
	                   nil
	   ===============================================
	*/
	//nil 讲解
	// 不同类型的默认值
	/*
	   bool     false
	   number   0
	   string    ""
	   point     nil
	   slice     nil
	   map       nil
	   channel   nil
	   interface nil
	   func      nil

	   struct    内部字段的默认值
	*/
	var o []int //nil slice
	if o == nil {
		fmt.Println("nil")
	}

	o = make([]int, 0) // empty slice
	if o != nil {
		fmt.Println("Not nil")
	}

	/*
	   =========================================================
	                         接口
	   =========================================================
	*/

	// 鸭子类型

	// 定义接口
	// type xxx interface{
	//
	//    method_xx()
	//    call() string
	// }

	var duck Duck = &donaldDuck{"donaldDuck"}
	duck.call()
	duck.walk()

	//多接口多结构体
	var wc1 MyWriter = &WriterCloser{}
	_ = wc1.Writer()

	var wc2 Closer = &WriterCloser{}
	_ = wc2.closer()

	wc1 = &WriterCloser{}
	_ = wc1.Writer()

	wc2 = &WriterCloser{}
	_ = wc2.closer()

	// 类型的断言
	a1 := 3
	b1 := 10

	r1 := add(a1, b1)
	fmt.Println(r1)

	var a2 int32 = 3
	var b2 int32 = 3100

	r2 := add(a2, b2)
	//r
	if r3, ok := r2.(int32); ok {
		fmt.Println(r3 + 10)
	}

	fmt.Println(r2)

	// 接口嵌套
	var rw readWriter = &MyreaderWriter{}
	rw.readWriter("")
	rw.Writer()
	rw.Reader()

}

func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case int32:
		return a.(int32) + b.(int32)
	default:
		panic("xxxx")
	}

	a1, _ := a.(int)
	b1, _ := b.(int)
	return a1 + b1
}

type MyreaderWriter struct {
}

func (m *MyreaderWriter) Reader() error {
	fmt.Println("Reader")
	return nil
}

func (m *MyreaderWriter) Writer() error {
	fmt.Println("Writer")
	return nil
}

func (m *MyreaderWriter) readWriter(str string) string {
	fmt.Println("readWriter")
	return str
}

type readWriter interface {
	MyReader
	MyWriter

	readWriter(string) string
}

type MyReader interface {
	Reader() error
}

type MyWriter interface {
	Writer() error
}

type Closer interface {
	closer() error
}

type WriterCloser struct {
	//嵌套
	MyWriter
}

func (wc *WriterCloser) Writer() error {
	fmt.Println("WriterCloser Writer")
	return nil
}

func (wc *WriterCloser) closer() error {
	fmt.Println("WriterCloser closer")
	return nil
}

type WriterCloser2 struct {
	Closer
}

func (wc *WriterCloser2) Writer() error {
	fmt.Println("WriterCloser2 Writer")
	return nil
}

func (wc *WriterCloser2) closer() error {
	fmt.Println("WriterCloser2 closer")
	return nil
}

type Duck interface {
	call()
	walk()
}

type donaldDuck struct {
	name string
}

func (d *donaldDuck) call() {
	fmt.Println(d.name)
}

func (d *donaldDuck) walk() {
	fmt.Println("walking")
}

//值传递需要修改变量的值，不能是指针变量的交换
//交换复制中的数据没用，应该去修改原始的数据内容
func swap(a, b *int) {
	//a,b = b,a

	temp := *a
	*a = *b
	*b = temp
}

//指针
//
func updateName(p *Person) {
	p.addr = "china"
}

func updateName2(p Person) {
	p.addr = "china"
}

type Person struct {
	name string
	age  int
	addr string
}

//板顶函数
func (p Person) toString() {
	p.name = "toString"
	fmt.Printf("+++++ name:%s,age:%d,addr:%s \n", p.name, p.age, p.addr)
}

//板顶函数 //数据很大不用copy,可以使用指针
func (p *Person) print() {
	p.name = "print"
	fmt.Printf("------name:%s,age:%d,addr:%s \n", p.name, p.age, p.addr)
}

type Student struct {
	//第一种嵌套（类似包装类）
	//p Person

	//第二种嵌套 匿名嵌入
	Person

	score float32

	//同名外层优先级高
	name string
}
