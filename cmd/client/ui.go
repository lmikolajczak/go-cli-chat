package main

import (
	"github.com/jroimartin/gocui"
	"golang.org/x/net/websocket"
)

type UI struct {
	*gocui.Gui
	Connection *websocket.Conn
}

func NewUI(connection *websocket.Conn) (*UI, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}
	ui := &UI{Gui: g, Connection: connection}

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

	if name, err := g.SetView("name", maxX/2-10, maxY/2-1, maxX/2+10, maxY/2+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.SetCurrentView("name")
		name.Title = "name"
		name.Autoscroll = false
		name.Wrap = true
		name.Editable = true
	}
	return nil
}

func (ui *UI) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (ui *UI) setName(g *gocui.Gui, v *gocui.View) error {
	g.SetViewOnBottom("name")
	g.SetCurrentView("input")
	return nil
}

func (ui *UI) sendMsg(g *gocui.Gui, v *gocui.View) error {
	v.SetCursor(0, 0)
	v.Clear()
	return nil
}
