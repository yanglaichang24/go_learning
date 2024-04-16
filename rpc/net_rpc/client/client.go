package main

import (
	"fmt"
	"net/rpc"
)

/**

rpc

*/
func main() {

	// rpc 连接服务器
	client, e := rpc.Dial("tcp", "localhost:8090")
	if e != nil {
		return
	}
	defer client.Close()

	var resp string

	//调用远程函数
	if e = client.Call("hello.HelloWord", "你好", &resp); e != nil {
		return
	}
	fmt.Println(resp)

}
