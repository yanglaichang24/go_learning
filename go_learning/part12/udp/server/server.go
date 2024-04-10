package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:50000")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	var buffer [512]byte
	for {
		n, addr, err := conn.ReadFromUDP(buffer[0:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Received data:", string(buffer[0:n]))
		n, err = conn.WriteToUDP([]byte("Hello, client!"), addr)
		if err != nil {
			fmt.Println(err)
		}
	}
}
