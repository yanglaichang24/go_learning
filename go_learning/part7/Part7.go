package main

import (
	"errors"
	"fmt"
)

/*
	  函数
	    -普通函数
		-匿名函数
	    -闭包

	  Go函数是"一等公民"
	     1.函数本身可以作为变量
	     2.匿名函数，闭包
	     3.函数可以满足接口
*/
func main() {
	a := 1
	if add, e := add(a, 3); e == nil {
		fmt.Println(add)
		fmt.Println(a)
	}

	if add, e := add4(a, 300); e == nil {
		fmt.Println(add)
		fmt.Println(a)
	}

	//函数本身可以作为变量
	f2 := add2
	if f21, e := f2(1, 2); e == nil {
		fmt.Println(f21)
	}

	i := calc("+")
	i()

	// 匿名函数
	i1 := func() int {
		return 100
	}

	i2 := calc1(i1)

	i3 := calc2(100, 100, func(a, b int) int {
		return a + b
	})

	i4 := calc2(100, 100, func(a, b int) int {
		return a - b
	})

	i5 := calc2(100, 100, func(a, b int) int {
		return a * b
	})

	i6 := calc2(100, 100, add3)

	fmt.Println(i2, i3, i4, i5, i6)

	// 闭包
	increment := autoIncrement()
	for l := 0; l < 10; l++ {
		fmt.Println(increment())
	}

	// defer 相当于finally
	// 最后关闭资源、锁，连接、文件
	deferTest()
	fmt.Println(deferTest1())

	//error  panic  recover
	if _, e := fun11(); e != nil {
		fmt.Println(e)
	}
	//panic 程序挂掉
	/*if _, e := fun12(); e != nil {
		fmt.Println(e)
	}*/

	//recover 能捕捉到panic
	fmt.Println("============================================")
	if r, e := fun13(); e != nil {
		fmt.Println("print error ", e)
	} else {
		fmt.Println("No error ", r)
	}

}

// recover
func fun13() (a int, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover.....", r)
			err = errors.New(" error ")
		}
	}()
	//启动检查
	//panic("this is panic")
	var m1 map[string]string
	m1["id"] = "m1"
	//panic: assignment to entry in nil map

	fmt.Println("panic post .....")

	return 1, errors.New("this is error")
}

// panic
func fun12() (int, error) {
	//启动检查
	//panic("this is panic")
	var m1 map[string]string
	m1["id"] = "m1"
	//panic: assignment to entry in nil map
	return 1, errors.New("this is error")
}

// error
func fun11() (int, error) {
	return 1, errors.New("this is error")
}

func deferTest1() (ret int) {
	defer func() {
		ret++
	}()
	return 10
}

func deferTest() {
	lock()
	defer fmt.Println("defer post fmt")
	defer unlock() // 会放在return之前执行的
	defer fmt.Println("defer pre fmt")

	fmt.Println("step 1")
	fmt.Println("step 2")
	fmt.Println("step 3")
}

func lock() {
	fmt.Println("locking")
}

func unlock() {
	fmt.Println("unlocking")
}

// 函数值传递
func add(a int, b int) (int, error) {
	c := a + b
	a += 100
	fmt.Println(a)
	return c, nil
}

func add2(a, b int) (int, error) {
	return a + b, nil
}

func add3(a, b int) int {
	return a + b
}

func fun4() {

}

func fun5() (sum int, e error) {
	return
}

func fun6(a ...int) (sum int, e error) {
	for i := 0; i < len(a); i++ {
		sum += i
	}
	return sum, e
}

func add4(a int, b int) (sum int, e error) {
	sum = a + b
	a += 100
	fmt.Println(a)
	return sum, e
}

func calc(op string) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("加法")
		}

	case "-":
		return func() {
			fmt.Println("加法")
		}
	default:
		return func() {
			fmt.Println("Not mattch")
		}
	}
}

func calc1(myfun func() int) int {
	return myfun()
}

func calc2(a, b int, myfun func(int, int) int) int {
	return myfun(a, b)
}

func calc3(a, b int, myfun func(c, d int) int) int {
	return myfun(a, b)
}

func autoIncrement() func() int {
	local := 0
	return func() int {
		local++
		return local
	}
}
