package main

import (
	"log"
	"runtime"
	"time"

	"github.com/Luqqk/go-cli-chat/internal/data"
)

func main() {
	chat := chat{
		users: make([]*data.User, 0),
		emit:  make(chan *data.Message),
		event: make(chan *data.Event),
	}
	go func() {
		for {
			<-time.After(3 * time.Second)
			log.Println("threads:", runtime.NumGoroutine(), "users:", len(chat.users))
		}
	}()
	chat.serve()
}
