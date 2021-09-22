package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	chat := chat{
		connections: make([]*websocket.Conn, 0),
		emit:        make(chan message),
		mutex:       sync.Mutex{},
	}
	go func() {
		for {
			<-time.After(3 * time.Second)
			log.Println("threads:", runtime.NumGoroutine(), "users:", len(chat.connections))
		}
	}()
	chat.serve()
}
