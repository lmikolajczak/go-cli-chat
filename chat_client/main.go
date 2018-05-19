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

	g.SetManagerFunc(Layout)

	g.SetKeybinding("name", gocui.KeyEnter, gocui.ModNone, Connect)

	g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, Send)

	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Disconnect)

	g.MainLoop()
}
