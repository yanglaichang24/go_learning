package main

import "fmt"

/*
  条件判断与for循环

  if <condition> {
     ****
  }


*/

func main() {

	// 条件判断
	age := 16
	addr := "china"

	if age == 16 && addr == "china" {
		fmt.Println(".....")
	}

	if (age == 16) && (addr == "china") {
		fmt.Println(".....")
	}

	if age == 16 {
		if addr == "china" {

		} else {

		}
	} else if age == 19 {

	}

	// for循环

	/*
	   for init;condition;post{

	    }
	*/

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
		break
	}

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
	fmt.Println("----------------------------------------")
	for v := range l {
		fmt.Println(v)
	}

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

	// goto 跳转到指定代码块执行
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break
			}
			fmt.Println(i, j)
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("-------------------------------------------------------")

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

	//switch
	/*
		  switch var {
		      case var2:
		            ....
		      case var3:
		            ....
		      default:
		            .....
		}
	*/

	day := "3"
	switch day {
	case "1":
		fmt.Println("wed")
	case "2":
		fmt.Println("Tud")
	default:
		fmt.Println("ext")
	}

	switch day {
	case "1", "4":
		fmt.Println("wed")
	case "2", "9":
		fmt.Println("Tud")
	default:
		fmt.Println("ext")
	}

	switch {
	case day == "1":
		fmt.Println("wed")
	case day == "2":
		fmt.Println("Tud")
	default:
		fmt.Println("ext")
	}

}
