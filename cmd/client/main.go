package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()
	g.SetManagerFunc(layout)
	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	g.MainLoop()
	// connection, err := websocket.Dial("ws://server:5000/", "", "http://server/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer connection.Close()
	// go func() {
	// 	for {
	// 		message := data.NewMessage()
	// 		err := websocket.JSON.Receive(connection, message)
	// 		if err != nil {
	// 			return
	// 		}
	// 		log.Println("message:", message)
	// 	}
	// }()
	// <-time.After(2 * time.Second)
	// message := data.NewMessage()
	// message.From = "Client"
	// message.Text = "test message from the client"
	// websocket.JSON.Send(connection, message)
	// <-time.After(2 * time.Second)
	// message.Text = "another message from the client"
	// websocket.JSON.Send(connection, message)
	// <-time.After(10 * time.Second)
}
