package main

import (
	"fmt"
	"net/rpc"
)

/**

rpc

*/

type Myclient struct {
	c *rpc.Client
}

func init_client(addr string) Myclient {
	client, e := rpc.Dial("tcp", addr)
	if e != nil {
		return
	}
	return Myclient{c: client}
}

func (mc *Myclient) HelloWord(str string, resp *string) error {
	return mc.c.Call("hello.HelloWord", "你好", &resp)
}

func (mc *Myclient) Close() error {
	return mc.c.Close()
}

/*
  封装接口
*/
func main() {
	// rpc 连接服务器
	client := init_client("localhost:8093")
	defer client.Close()

	var resp *string
	_ := client.HelloWord("xxxx", resp)
	fmt.Println("resp")
}
