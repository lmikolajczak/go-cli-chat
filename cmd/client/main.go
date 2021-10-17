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
	defer ui.Close()

	ui.SetManagerFunc(ui.layout)
	if err := ui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, ui.quit); err != nil {
		log.Fatalln(err)
	}
	if err := ui.SetKeybinding("name", gocui.KeyEnter, gocui.ModNone, ui.setName); err != nil {
		log.Fatalln(err)
	}
	if err := ui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, ui.sendMsg); err != nil {
		log.Fatalln(err)
	}
	go ui.receiveMsg()
	if err = ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}
