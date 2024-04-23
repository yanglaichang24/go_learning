Go [HOME ](../learning-go.md)
---
# 条件与循环

## 条件

Go 语言提供了以下几种条件判断语句：

| 语句           | 	描述                                                                          |
|--------------|------------------------------------------------------------------------------|
| if 语句	       | if 语句 由一个布尔表达式后紧跟一个或多个语句组成。                                                  |
| if...else 语句 | 	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。                      |
| if 嵌套语句	     | 你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。                                |
| switch 语句	   | switch 语句用于基于不同条件执行不同动作。                                                     |
| select 语句    | 	select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。 |

### if
 ``` go
  if <condition> {
     ****
  }
  
    age := 16
	addr := "china"

	if age == 16 && addr == "china" {
		fmt.Println(".....")
	}
 ```

### if-else

```go
    addr := "china"
	
    if addr == "china" {
        fmt.Println("----")
    } else {
        fmt.Println("@@@")
    }
```

### if 嵌套
  ```go
    age := 16
    addr := "china"
    if age == 16 {
        if addr == "china" {

        } else {

        }
    } else if age == 19 {

    }
  ```
### switch 
```go
    day := "3"
	switch day {
        case "1":
            fmt.Println("wed")
        case "2":
            fmt.Println("Tud")
        default:
            fmt.Println("ext")
    }
```

### select 

```go
    ch11 := make(chan struct{})
    ch21 := make(chan struct{})

    select {
        case <-ch11:
            fmt.Println("ch11 down")
        case <-ch21:
            fmt.Println("ch21 down ")
    }
	
```

## 循环

Go 语言提供了以下几种类型循环处理语句：

| 循环类型    | 	描述                     |
|---------|-------------------------|
| for 循环	 | 重复执行语句块                 |
| 循环嵌套	   | 在 for 循环中嵌套一个或多个 for 循环 |

循环控制语句
循环控制语句可以控制循环体内语句的执行过程。

GO 语言支持以下几种循环控制语句：

| 控制语句        | 	描述                           |
|-------------|-------------------------------|
| break 语句    | 	经常用于中断当前 for 循环或跳出 switch 语句 |
| continue 语句 | 	跳过当前循环的剩余语句，然后继续进行下一轮循环。     |
| goto 语句	    | 将控制转移到被标记的语句。                 |


### for

```go
       // for init;condition;post{
       // }

        for i := 1; i < 10; i++ {
            fmt.Println(i)
        }
        j := 0
        for ; j < 10; j++ {
            fmt.Println(j)
        }

        for j < 11 {
            j++
            fmt.Println(j)
		}

        for {
            fmt.Println("---")
        }
```
### for range 
```go
    // for range 针对字符串、数组，切片，map.channel

	/*
	   for key ,value := range xx{}
	*/

        l := "我是中国人abc，我爱祖国a"
        for index, v := range l {
            fmt.Printf("%d,%c \n", index, v)
        }
        
        for _, v := range l {
            fmt.Printf("%c \n", v)
        }
        
        // 注意
        for v := range l {
            fmt.Println(v)
        }

```

### for 嵌套
```go
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 2 {
                break
            }
            fmt.Println(i, j)
        }
    }
	
```

### continue/break

```go
        //退出条件 continue/break
    i := 0
    for {
        i++
        if i == 5 {
            continue
        }
        if i > 10 {
            break
        }
        fmt.Println(i)
    }
```
### goto

```go
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 2 {
                goto over
            }
	    fmt.Println(i, j)
         }
    }
    
    over:
    fmt.Println("xxxxxx go")

```





---
[PRE]()    
NEXT >> [NEXT]()