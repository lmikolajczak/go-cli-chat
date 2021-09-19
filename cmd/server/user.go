package main

import (
	"bufio"
	"fmt"
	"net"
)

type User struct {
	connection net.Conn
	buffer     *bufio.ReadWriter
	read       chan string
	write      chan string
}

func NewUser(connection net.Conn) *User {
	rw := bufio.NewReadWriter(
		bufio.NewReader(connection),
		bufio.NewWriter(connection),
	)

	user := &User{
		connection: connection,
		buffer:     rw,
		read:       make(chan string),
		write:      make(chan string),
	}

	go user.Read()
	go user.Write()

	return user
}

func (u *User) Read() {
	for {
		message, err := u.buffer.ReadString('\n')
		if err != nil {
			// As soon as we receive EOF (connection has been closed)
			// we should disconnect the user and cleanup
			fmt.Println(err)
			u.Disconnect()
			return
		}
		fmt.Println("New message from client")
		u.read <- message
	}
}

func (u *User) Write() {
	for msg := range u.write {
		_, err := u.buffer.WriteString(msg)
		if err != nil {
			return
		}
	}
}

func (u *User) Disconnect() {
	close(u.read)
	close(u.write)
}
