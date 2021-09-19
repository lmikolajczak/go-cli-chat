package main

import (
	"fmt"
	"net"
	"runtime"
	"sync"
	"time"
)

type server struct {
	users []*User
	emit  chan string
	mu    sync.Mutex
}

func (s *server) Listen() error {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		return err
	}

	go func() {
		for {
			<-time.After(3 * time.Second)
			fmt.Println(s.users, runtime.NumGoroutine())
		}
	}()

	go s.broadcast()

	for {
		connection, err := listener.Accept()
		if err != nil {
			return err
		}
		s.join(connection)
	}
}

func (s *server) join(connection net.Conn) {
	user := NewUser(connection)
	s.users = append(s.users, user)

	go func() {
		defer s.remove(user)
		for msg := range user.read {
			fmt.Println("Push to emit")
			s.emit <- msg
		}
	}()
}

func (s *server) broadcast() {
	for msg := range s.emit {
		fmt.Println("Broadcast", msg)
	}
}

func (s *server) remove(user *User) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := len(s.users) - 1; i >= 0; i-- {
		if s.users[i] == user {
			s.users = append(s.users[:i], s.users[i+1:]...)
		}
	}
}
