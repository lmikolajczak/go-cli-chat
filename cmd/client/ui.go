package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lmikolajczak/go-cli-chat/internal/data"
	"golang.org/x/net/websocket"
)

type UI struct {
	*gocui.Gui
	Username   string
	Connection *websocket.Conn
}

func NewUI(connection *websocket.Conn, username string) (*UI, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}
	ui := &UI{Gui: g, Connection: connection, Username: username}

	return ui, nil
}

func (ui *UI) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	g.Cursor = true

	if messages, err := g.SetView("messages", 0, 0, maxX-20, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		messages.Title = "messages"
		messages.Autoscroll = true
		messages.Wrap = true
	}

	if input, err := g.SetView("input", 0, maxY-5, maxX-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		input.Title = "send"
		input.Autoscroll = false
		input.Wrap = true
		input.Editable = true
	}

	if users, err := g.SetView("users", maxX-20, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		users.Title = "users"
		users.Autoscroll = false
		users.Wrap = true
	}
	g.SetCurrentView("input")

	return nil
}

func (ui *UI) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (ui *UI) sendMsg(g *gocui.Gui, v *gocui.View) error {
	if len(v.Buffer()) == 0 {
		v.SetCursor(0, 0)
		v.Clear()
		return nil
	}
	message := data.Message{
		From: ui.Username,
		Text: v.Buffer(),
	}
	err := websocket.JSON.Send(ui.Connection, message)
	if err != nil {
		return err
	}
	v.SetCursor(0, 0)
	v.Clear()
	return nil
}

func (ui *UI) receiveMsg() {
	for {
		payload := &data.Payload{}
		err := websocket.JSON.Receive(ui.Connection, payload)
		if err != nil {
			return
		}

		ui.Update(func(g *gocui.Gui) error {
			switch {
			case payload.Message != data.Message{}:
				view, _ := ui.View("messages")
				fmt.Fprint(view, payload.Message.Formatted())
			case payload.Notification != data.Notification{}:
				view, _ := ui.View("users")
				switch payload.Notification.Type {
				case data.ConnectedUsers:
					view.Clear()
					fmt.Fprint(view, payload.Notification.Formatted())
				}
			}
			return nil
		})
	}
}
