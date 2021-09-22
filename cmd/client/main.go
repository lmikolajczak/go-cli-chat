package main

import (
	"log"
	"time"

	"golang.org/x/net/websocket"
)

type message struct {
	Text string `json:"text"`
}

func main() {
	connection, err := websocket.Dial("ws://server:5000/", "", "http://server/")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	go func() {
		for {
			message := message{}
			err := websocket.JSON.Receive(connection, &message)
			if err != nil {
				return
			}
			log.Println("message:", message)
		}
	}()
	<-time.After(2 * time.Second)
	message := message{"test message from the client"}
	websocket.JSON.Send(connection, message)
	<-time.After(2 * time.Second)
	message.Text = "another message from the client"
	websocket.JSON.Send(connection, message)
	<-time.After(10 * time.Second)
}
