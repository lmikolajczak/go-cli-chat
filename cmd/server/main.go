package main

import (
	"log"
	"sync"
)

func main() {
	server := server{
		users: []*User{},
		emit:  make(chan string),
		mu:    sync.Mutex{},
	}
	err := server.Listen()
	if err != nil {
		log.Fatal(err)
	}
}
