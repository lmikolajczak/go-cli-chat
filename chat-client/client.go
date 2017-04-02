package main

import (
	"go-cli-chat/chat-client/ui"
	"go-cli-chat/chat-client/utils"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	g.SetManagerFunc(ui.Layout)

	g.SetKeybinding("nick", gocui.KeyEnter, gocui.ModNone, utils.Connect)

	g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, utils.Send)

	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, utils.Disconnect)

	g.MainLoop()
}
