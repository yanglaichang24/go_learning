package main

import (
	"context"
	"fmt"
	common "go_pro/go_learning"
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

	common.PrintLine("  sync.WaitGroup ")

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

	common.PrintLine(" lock 锁")

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
	common.PrintLine("读写锁")

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

	//goroutine 通信
	common.PrintLine("goroutine 相互通信 channel")
	// channel 是线程安全的
	//channel定义//通道里面放什么
	var msg chan string

	// 有缓冲区和无缓冲区
	msg = make(chan string, 1) //channel 初始化为0 永远阻塞，请注意

	//放入channel
	msg <- "msg"
	//从channel读取
	data := <-msg
	fmt.Println(data)

	//channel 初始化为0 永远阻塞
	//msg := make(chan string, 0)
	//msg <- "msg"
	//data:= <- msg

	msg0 := make(chan string, 0) //channel 初始化为0 永远阻塞，除非在goroutine
	go func(msg chan string) {   //主要是happen-before保证可见性
		data := <-msg
		fmt.Println(data)
	}(msg0)

	msg0 <- "msg"

	/*
	   有缓冲与无缓冲channel
	   无缓冲channel A 发生 B 立马需要知道
	   有缓冲channel 缓存一部分数据

	   go 中channel的应用
	         消息传递与消息过滤
	         信号广播
	         事件订阅和广播
	         任务分发
	         结果汇总
	         并发控制
	         同步与异步
	         。。。。。
	*/

	// for range 获取
	// close 用法

	d := make(chan int, 2)

	// 读取
	go func(msg chan int) {
		for data := range msg {
			fmt.Println(data)
		}
		////如果不关闭，这里永远执行不到
		fmt.Println("结束了不获取了")
	}(d)

	// 写入
	d <- 1
	d <- 2
	d <- 3

	//关闭channel
	close(d)
	//关闭的channel 可以读不可以写

	time.Sleep(20000)

	// 默认情况下channel 是全双工的，有时候需要单向channel
	var _ chan int   //双向
	var _ chan<- int //单项写入
	var _ <-chan int //单项读取

	d3 := make(chan int, 3)
	var send chan<- int = d3 //read-only //不能逆向处理
	var rec <-chan int = d3  // rec-only  //不能逆向处理

	send <- 2
	// <-send 无法用
	<-rec

	d6 := make(chan int, 10)
	go produce(d6)
	go comsuer(d6)

	time.Sleep(10 * time.Second)

	//交叉打印
	//自行解决

	// select 作用于channel,类似linux select/poll/epoll
	common.PrintLine(" select ")

	//监控多个任务其中一个完成就完成
	go g1()
	go g2()

	// 阻塞等待读取,channel 是线程安全的
	<-ch
	fmt.Println("完成")

	//监控多个任务都完成在完成
	//多个channel处理
	ch11 := make(chan struct{})
	ch21 := make(chan struct{})

	go g11(ch11)
	go g21(ch21)

	// 谁就绪就处理谁，同时就绪就是随机（防止饥饿）

	select {
	case <-ch11:
		fmt.Println("ch11 down")
	case <-ch21:
		fmt.Println("ch21 down ")
	}

	//超时机制
	// (方法一)
	/*for {
		select {
		case <-ch11:
			fmt.Println("ch11 down")
		case <-ch21:
			fmt.Println("ch21 down ")
		default: //没有就绪
			time.Sleep(10 * time.Second)
			fmt.Println("xxx")
		}
	}*/

	// 特殊的超时channle(推荐)
	/*timer := time.NewTimer(5 * time.Second)
	select {
	case <-ch11:
		fmt.Println("ch11 down")
	case <-ch21:
		fmt.Println("ch21 down ")
	case <-timer.C:
		fmt.Println(" ")
	}*/

	common.PrintLine(" context ")

	//主线程控制子线程介素
	//方法一：可以利用上面的无缓存的select/channel进行处理
	//方法二：context // withCancel

	wg20.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	go getCpu(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	wg20.Wait()
	fmt.Println("任务完成 ")

	//context 层级控制/父类cancel/子类依然有效果
	/*
	  ctx, cancel := context.WithCancel(context.Background())
	  ctx2,_ := context.WithCancel(ctx)
	  go getCpu(ctx2)
	  time.Sleep(3 * time.Second)
	  cancel()
	  wg20.Wait()

	   cancel ------>  cancel ---->  cancel
	  --------        --------      --------
	  -  ctx    - --- - ctx1 - ---- - ctx2 -
	  --------        --------      --------

	*/

	// context withTimeout/context Deadline
	// 主动超时
	// ctx, cancel := context.WithTimeout(context.Background(),3 * time.Second)
	// go getCpu(ctx)
	// wg20.Wait()

	// context withValue 传值
	//value := context.WithValue(context.Background(), "id", "1")

	// ctx.Value("id")

	//TODO
	//退出return/exit/goexit

	// os.Exit(-1)
	// runtime.Goexit()

}

var wg20 sync.WaitGroup

func getCpu(ctx context.Context) {

	defer wg20.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("结束了")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("打印cpu 消息")
		}
	}

}

//空的结构体
var ch = make(chan struct{})

func g1() {
	fmt.Println("g1 done ")
	ch <- struct{}{}

}

func g11(ch11 chan struct{}) {
	fmt.Println("g11 done ")
	ch11 <- struct{}{}

}

func g2() {
	fmt.Println("g2 done ")
	ch <- struct{}{}
}

func g21(ch21 chan struct{}) {
	fmt.Println("g21 done ")
	ch21 <- struct{}{}
}

func produce(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func comsuer(c <-chan int) {
	for num := range c {
		fmt.Println(num)
	}

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
