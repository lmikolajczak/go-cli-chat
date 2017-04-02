package main

import (
	"fmt"
	"go-cli-chat/chat-server/chat"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}

	chat := chat.CreateChat()

	for { // listen for connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		chat.Connect(conn)
	}
}
