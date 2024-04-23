Go [HOME ](../learning-go.md)
---
# 数据结构实现

## 数组

### 声明数组
Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：

> var arrayName [size]dataType

其中，arrayName 是数组的名称，size 是数组的大小，dataType 是数组中元素的数据类型。

### 初始化数组
以下演示了数组初始化：
```go
  //以下实例声明一个名为 numbers 的整数数组，其大小为 5，
  //在声明时，数组中的每个元素都会根据其数据类型进行默认初始化，
  //对于整数类型，初始值为 0。
  
    var numbers [5]int

 //还可以使用初始化列表来初始化数组的元素：

  var numbers = [5]int{1, 2, 3, 4, 5}
  
//以上代码声明一个大小为 5 的整数数组，
//并将其中的元素分别初始化为 1、2、3、4 和 5。

//另外，还可以使用 := 简短声明语法来声明和初始化数组：

numbers := [5]int{1, 2, 3, 4, 5}

//以上代码创建一个名为 numbers 的整数数组，
//并将其大小设置为 5，并初始化元素的值。

//注意：在 Go 语言中，数组的大小是类型的一部分，
//因此不同大小的数组是不兼容的，
//也就是说 [5]int 和 [10]int 是不同的类型。

//以下定义了数组 balance 长度为 5 类型为 float32，
//并初始化数组的元素：

var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

//我们也可以通过字面量在声明数组的同时快速初始化数组：

balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

//如果数组长度不确定，可以使用 ... 代替数组的长度，
//编译器会根据元素个数自行推断数组的长度：

var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

//或

balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

//如果设置了数组的长度，我们还可以通过指定下标来初始化元素：
//将索引为 1 和 3 的元素初始化
balance := [5]float32{1:2.0,3:7.0}

//初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

//如果忽略 [] 中的数字不设置数组大小，
//Go 语言会根据元素的个数来设置数组的大小：

balance[4] = 50.0

//以上实例读取了第五个元素。
//数组元素可以通过索引（位置）来读取（或者修改），
//索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。

```

### 访问数组元素
数组元素可以通过索引（位置）来读取。
格式为数组名后加中括号，中括号中为索引的值。例如：

> var salary float32 = balance[9]

以上实例读取了数组 balance 第 10 个元素的值。

### 遍历
```go
    arr3 := [3]string{"a", "b", "c"}
    for key, value := range arr3 {
        fmt.Println(key, value)
    }
```
### 比较

```go

    //数组比较 类型一样
    //if arr2 == arr3{} //
    if arr == arr3 {
        fmt.Println("=")
    } else {
        fmt.Println("!=")
    }
```
### 多维数组
```go
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
}
	
```
## 切片

Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

### 定义切片

你可以声明一个未指定大小的数组来定义切片：

> var identifier []type

切片不需要说明长度。

或使用 make() 函数来创建切片:

> var slice1 []type = make([]type, len)

也可以简写为

> slice1 := make([]type, len)
> 
也可以指定容量，其中 capacity 为可选参数。

> make([]T, length, capacity)
> 
这里 len 是数组的长度并且也是切片的初始长度。

### 切片初始化

> s :=[] int {1,2,3 }

直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。

> s := arr[:]

初始化切片 s，是数组 arr 的引用。

> s := arr[startIndex:endIndex]

将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。

> s := arr[startIndex:]

默认 endIndex 时将表示一直到arr的最后一个元素。

> s := arr[:endIndex]

默认 startIndex 时将表示从 arr 的第一个元素开始。

> s1 := s[startIndex:endIndex]

通过切片 s 初始化切片 s1。

> s :=make([]int,len,cap)

通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。

### len() 和 cap() 函数

切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

以下为具体实例：

实例
```go

package main

import "fmt"

func main() {
var numbers = make([]int,3,5)
    printSlice(numbers)
}

func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
//以上实例运行输出结果为:
//len=3 cap=5 slice=[0 0 0]
```
### 空(nil)切片

一个切片在未初始化之前默认为 nil，长度为 0

### 切片截取
可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]，实例如下：

```go

package main
import "fmt"

func main() {
    /* 创建切片 */
    numbers := []int{0,1,2,3,4,5,6,7,8}  
    printSlice(numbers)

    /* 打印原始切片 */
    fmt.Println("numbers ==", numbers)

    /* 打印子切片从索引1(包含) 到索引4(不包含)*/
    fmt.Println("numbers[1:4] ==", numbers[1:4])

    /* 默认下限为 0*/
    fmt.Println("numbers[:3] ==", numbers[:3])

    /* 默认上限为 len(s)*/
    fmt.Println("numbers[4:] ==", numbers[4:])

    numbers1 := make([]int,0,5)
    printSlice(numbers1)

    /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
    number2 := numbers[:2]
    printSlice(number2)

    /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
    number3 := numbers[2:5]
    printSlice(number3)

}

func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

```

### append() 和 copy() 函数
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。

实例
package main

import "fmt"

func main() {
var numbers []int
printSlice(numbers)

/* 允许追加空切片 */
numbers = append(numbers, 0)
printSlice(numbers)

/* 向切片添加一个元素 */
numbers = append(numbers, 1)
printSlice(numbers)

/* 同时添加多个元素 */
numbers = append(numbers, 2,3,4)
printSlice(numbers)

/* 创建切片 numbers1 是之前切片的两倍容量*/
numbers1 := make([]int, len(numbers), (cap(numbers))*2)

/* 拷贝 numbers 的内容到 numbers1 */
copy(numbers1,numbers)
printSlice(numbers1)  
}

func printSlice(x []int){
fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
以上代码执行输出结果为：

len=0 cap=0 slice=[]
len=1 cap=1 slice=[0]
len=2 cap=2 slice=[0 1]
len=5 cap=6 slice=[0 1 2 3 4]
len=5 cap=12 slice=[0 1 2 3 4]


## Map
## List


---
[PRE]()    
NEXT >> [NEXT]()