package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloWord struct {
}

func (hw *HelloWord) HelloWord(name string, resp *string) error {
	*resp = name + "你好"
	return nil
}

/**

rpc包提供了通过网络或其他I/O连接对一个对象的导出方法的访问。服务端注册一个对象，使它作为一个服务被暴露，服务的名字是该对象的类型名。注册之后，对象的导出方法就可以被远程访问。
服务端可以注册多个不同类型的对象（服务），但注册具有相同类型的多个对象是错误的

*/
func main() {

	//注册服务
	if err := rpc.RegisterName("hello", new(HelloWord)); err != nil {
		fmt.Println("注册失败")
		return
	}

	// 设置监听
	listener, e := net.Listen("tcp", "localhost:8090")
	if e != nil {
		fmt.Println("注册失败")
		return
	}

	defer listener.Close()

	fmt.Println("服务器开启")

	//建立连接
	conn, e := listener.Accept()
	if e != nil {
		fmt.Println("注册失败")
		return
	}

	//绑定服务
	rpc.ServeConn(conn)

}
