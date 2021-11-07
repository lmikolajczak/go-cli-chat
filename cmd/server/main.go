package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/Luqqk/go-cli-chat/internal/data"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 5000, "Server port")
	flag.Parse()
	chat := chat{
		users: make([]*data.User, 0),
		emit:  make(chan *data.Message),
		event: make(chan *data.Event),
	}
	// Report some basic server stats each 10 seconds
	go func() {
		for {
			time.Sleep(10 * time.Second)
			log.Println("threads:", runtime.NumGoroutine(), "users:", len(chat.users))
		}
	}()
	log.Println(fmt.Sprintf("Starting chat server at port %v...", port))
	chat.serve(port)
}
