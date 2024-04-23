package main

import (
	"fmt"
	"go_learning/part9/bean"
	// 别名用处
	v1 "go_learning/part9/bean/v1"
	//. "go_pro/go_learning/part9/bean/v1" 导入本地，不用写前缀报名，重名可读性差
	//_ "go_pro/go_learning/part9/bean/v1" 匿名 导入不用，特殊的函数init初始化
)

/*

 package /go module / go get

  go 语言中根据字母的大小写来确定访问的权限问题，如果是大写，作用域则可以被其他包访问
  如果首字母小写，作用域只能在本包中使用，包括接口，类型，函数，变量等

   go.mod文件

   go 命令
	   go install

	   go get xxxx下载
	   go get xxxx@version
	   go get -u xxx 升级
	   go get -u=patch

	   go mod tidy  //整理依赖，不需要的删除

	   GOPROXY 设置
	   go env -w GO111MODULE=on
	   go env -w GOPROXY=XXX

	   replace 使用 goMudule 替换

   代码规范
     命名规范
        包名
        变量名
        结构体
        接口命名
        常量命名




*/

func main() {

	user := bean.User{Name: "new "}
	user.Working()
	fmt.Println(user)

	u2 := bean.Default()
	u2.Working()
	fmt.Println(u2)

	//使用别名
	u3 := v1.User{Name: "v1 "}
	u3.Working()
	fmt.Println(u3)
}
