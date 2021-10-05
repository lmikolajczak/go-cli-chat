package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"golang.org/x/net/websocket"
)

func main() {
	connection, err := websocket.Dial("ws://server:5000/", "", "http://server/")
	if err != nil {
		log.Fatal(err)
	}
	ui, err := NewUI(connection)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Gui.Close()

	if err = ui.Gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
