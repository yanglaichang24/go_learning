package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Listening on localhost:50000...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	var buffer [512]byte
	for {
		n, err := conn.Read(buffer[0:])
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Println("Received data:", string(buffer[0:n]))
	}
}
