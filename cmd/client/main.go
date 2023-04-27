package main

import (
	"log"
	"os"

	"github.com/lmikolajczak/go-cli-chat/pkg/frontend"
)

func main() {
	ui, err := frontend.NewUI()
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	username := os.Args[1]
	ui.SetUsername(username)

	if err = ui.Connect(username); err != nil {
		log.Fatal(err)
	}

	ui.SetManagerFunc(ui.Layout)
	ui.SetKeyBindings(ui.Gui)

	if err = ui.Serve(); err != nil {
		log.Fatal(err)
	}
}
