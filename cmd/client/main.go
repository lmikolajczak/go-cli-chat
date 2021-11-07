package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
	"golang.org/x/net/websocket"
)

func main() {
	var username string
	var port int
	flag.StringVar(&username, "username", "", "Chat username")
	flag.IntVar(&port, "port", 5000, "Server port")
	flag.Parse()
	if len(strings.TrimSpace(username)) == 0 {
		log.Fatal("missing -username option (required)")
	}
	config, err := websocket.NewConfig(fmt.Sprintf("ws://:%v/", port), "http://")
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
	if err := ui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, ui.sendMsg); err != nil {
		log.Fatalln(err)
	}
	go ui.receiveMsg()
	if err = ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}
