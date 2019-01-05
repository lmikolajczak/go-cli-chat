package main

import (
	"fmt"
	"net"

	"github.com/Luqqk/go-cli-chat/pkg/server"
)

func main() {
	fmt.Println("Starting listening for connections...")

	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}

	chat := server.CreateChat()

	for { // listen for connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		chat.Connect(conn)
	}
}
