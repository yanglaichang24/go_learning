# Golang

Go （又称Golang）是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言

Go是从2007年末由Robert Griesemer, Rob Pike, Ken Thompson主持开发，后来还加入了Ian Lance Taylor, Russ Cox等人，并最终于2009年11月开源，在2012年早些时候发布了Go 1稳定版本。现在Go的开发已经是完全开放的，并且拥有一个活跃的社区。

## Go 语言特色
+ 简洁、快速、安全
+ 并行、有趣、开源
+ 内存管理、数组安全、编译迅速

## Go 语言用途
Go 语言被设计成一门应用于搭载 Web 服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。

对于高性能分布式系统领域而言，Go 语言无疑比大多数其它语言有着更高的开发效率。它提供了海量并行的支持，这对于游戏服务端的开发而言是再好不过了。

### 第一个 Go 程序
接下来我们来编写第一个 Go 程序 hello.go（Go 语言源文件的扩展是 .go），代码如下：
hello.go 文件

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

运行实例 
>要执行 Go 语言代码可以使用 go run 命令。

执行以上代码输出:
```
$ go run hello.go 
Hello, World!
```

此外我们还可以使用 go build 命令来生成二进制文件：
```
$ go build hello.go 
$ ls
hello    hello.go
$ ./hello 
Hello, World!

```

---

<h1> 学习大纲 </h1>

+ [基本语法结构/变量/常量](part1/part1.md)
+ [基础数据类型](part2/part.md)
+ [运算符](part3/part.md)
+ [字符串操作](part4/part.md)
+ [条件判断与循环](part5/part.md)
+ [集合类型](part6/part.md)
+ [函数](part7/part.md)
+ [结构体](part8/part.md)
+ [包/模块](part9/part.md)
+ [并发编程](part10/part.md)
+ [测试用例](part11/part.md)
+ [IO](part12/part.md)
+ [泛型](part13/part.md)



