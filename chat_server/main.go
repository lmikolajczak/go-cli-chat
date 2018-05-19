package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}

	chat := CreateChat()

	for { // listen for connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		chat.Connect(conn)
	}
}
