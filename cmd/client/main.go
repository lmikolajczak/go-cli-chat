package main

import (
	"flag"
	"log"

	"github.com/jroimartin/gocui"
	"golang.org/x/net/websocket"
)

func main() {
	var username string
	flag.StringVar(&username, "username", "anonymous", "Chat username")
	flag.Parse()
	config, err := websocket.NewConfig("ws://server:5000/", "http://server/")
	config.Header.Set("Username", username)
	if err != nil {
		log.Fatal(err)
	}
	connection, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	ui, err := NewUI(connection, username)
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
