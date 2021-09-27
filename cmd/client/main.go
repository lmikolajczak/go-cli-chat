package main

import (
	"log"
	"time"

	"github.com/Luqqk/go-cli-chat/internal/data"
	"golang.org/x/net/websocket"
)

func main() {
	connection, err := websocket.Dial("ws://server:5000/", "", "http://server/")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	go func() {
		for {
			message := data.NewMessage()
			err := websocket.JSON.Receive(connection, message)
			if err != nil {
				return
			}
			log.Println("message:", message)
		}
	}()
	<-time.After(2 * time.Second)
	message := data.NewMessage()
	message.From = "Client"
	message.Text = "test message from the client"
	websocket.JSON.Send(connection, message)
	<-time.After(2 * time.Second)
	message.Text = "another message from the client"
	websocket.JSON.Send(connection, message)
	<-time.After(10 * time.Second)
}
