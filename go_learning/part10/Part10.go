package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

/*
  并发编程

  用户线程/轻量级线程/协程/
  在用户空间需要实现协程的调度(GMP goroutine,M 线程 P 调度器) 类似线程池的概念（N:M）

  https://blog.csdn.net/sdasdas12/article/details/133170897

  相对于线程/进程都是操作系统内核空间调度

  内存占用小/切换快

*/

func printMsg() {
	fmt.Println("xxxx")
}

func main() {

	go printMsg()
	fmt.Println("main")
	time.Sleep(2 * time.Second)
	fmt.Println("post")

	go func() {
		fmt.Println("匿名函数")
	}()

	time.Sleep(2 * time.Second)

	//多个协程
	// 打印有重复//闭包
	/*for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}*/

	/*for i := 0; i < 10; i++ {
		//
		temp := i
		go func() {
			fmt.Println(temp)
		}()
	}*/

	/*for i := 0; i < 10; i++ {
		go func(t int) {
			fmt.Println(t)
		}(i)
	}*/

	/*
	  ==============================================
	                sync.WaitGroup
	  ==============================================
	*/

	// gotoutine 如何通知主goroutine结束，主goroutine 如何知道其他goroutine 已经结束
	fmt.Println("-------------")

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(t int) {
			defer wg.Done()
			fmt.Println(t)
		}(i)
	}

	wg.Wait()
	fmt.Println("all down")

	/*
	  ==============================================
	                lock
	  ==============================================
	*/
	wg1.Add(2)
	go add()
	go sub()

	wg1.Wait()
	fmt.Println("no safe =" + strconv.Itoa(total))

	// sync.Mutex 锁 与 并发自增的工具 atomic.AddInt32(&total2,1)
	// 查看相关代码

	//sync.Mutex
	wg2.Add(2)
	go safeAdd()
	go safeSub()

	wg2.Wait()
	fmt.Println("safe1 =" + strconv.Itoa(total1))

	//atomic.AddInt32
	wg3.Add(2)
	go safeAdd2()
	go safeSub2()

	wg3.Wait()
	fmt.Println("safe2 =" + strconv.Itoa(int(total2)))

	//读写锁
	fmt.Println(" =========  读写锁  =================== ")

	var wg5 sync.WaitGroup
	var rwlock sync.RWMutex
	var flag int

	wg5.Add(2)
	go func() {
		defer wg5.Done()
		time.Sleep(1000 * 10)
		rwlock.Lock()
		defer rwlock.Unlock()
		fmt.Println("获取写锁")
		for i := 0; i < 10; i++ {
			fmt.Println("处理步骤=" + strconv.Itoa(i))
			time.Sleep(1000)
		}
		flag = 12
		fmt.Println("处理完成")
	}()

	go func() {
		defer wg5.Done()
		var i int
		for {
			rwlock.RLock()
			fmt.Println("read flag = " + strconv.Itoa(flag))
			time.Sleep(10)
			rwlock.RUnlock()

			if i > 10 {
				break
			}
			i++
		}
	}()

	wg5.Wait()
	fmt.Println("done...")

	fmt.Println(" =========  goroutine 通信  =================== ")

}

var (
	wg1   sync.WaitGroup
	total int

	wg2    sync.WaitGroup
	lock   sync.Mutex
	total1 int

	wg3    sync.WaitGroup
	total2 int32
)

func add() {
	defer wg1.Done()
	for i := 0; i < 1000000; i++ {
		total += 1
	}
}

func safeAdd() {
	defer wg2.Done()
	for i := 0; i < 1000000; i++ {
		lock.Lock()
		total1 += 1
		lock.Unlock()
	}
}

func safeAdd2() {
	defer wg3.Done()
	for i := 0; i < 1000000; i++ {
		atomic.AddInt32(&total2, 1)
	}
}

func sub() {
	defer wg1.Done()
	for i := 0; i < 1000000; i++ {
		total -= 1
	}
}

func safeSub() {
	defer wg2.Done()
	for i := 0; i < 1000000; i++ {
		lock.Lock()
		total1 -= 1
		lock.Unlock()
	}
}

func safeSub2() {
	defer wg3.Done()
	for i := 0; i < 1000000; i++ {
		atomic.AddInt32(&total2, -1)
	}

}
