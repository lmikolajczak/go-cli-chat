package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/Luqqk/go-cli-chat/internal/data"
	"golang.org/x/net/websocket"
)

func main() {
	chat := chat{
		connections: make([]*websocket.Conn, 0),
		emit:        make(chan *data.Message),
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
