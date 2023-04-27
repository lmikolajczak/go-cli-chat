package frontend

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"golang.org/x/net/websocket"

	"github.com/lmikolajczak/go-cli-chat/pkg/chat"
)

const (
	WebsocketEndpoint = "ws://:3000/"
	WebsocketOrigin   = "http://"

	MessageWidget = "messages"
	UsersWidget   = "users"
	InputWidget   = "send"
)

type UI struct {
	*gocui.Gui

	username   string
	connection *websocket.Conn
}

func NewUI() (*UI, error) {
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, fmt.Errorf("NewUI: %w", err)
	}

	return &UI{Gui: gui}, nil
}

func (ui *UI) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	g.Cursor = true

	if messages, err := g.SetView(MessageWidget, 0, 0, maxX-20, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		messages.Title = MessageWidget
		messages.Autoscroll = true
		messages.Wrap = true
	}

	if input, err := g.SetView(InputWidget, 0, maxY-5, maxX-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		input.Title = InputWidget
		input.Autoscroll = false
		input.Wrap = true
		input.Editable = true
	}

	if users, err := g.SetView(UsersWidget, maxX-20, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		users.Title = UsersWidget
		users.Autoscroll = false
		users.Wrap = true
	}

	g.SetCurrentView(InputWidget)

	return nil
}

func (ui *UI) SetKeyBindings(g *gocui.Gui) error {
	if err := g.SetKeybinding(InputWidget, gocui.KeyCtrlC, gocui.ModNone, ui.Quit); err != nil {
		return err
	}

	if err := g.SetKeybinding(InputWidget, gocui.KeyEnter, gocui.ModNone, ui.WriteMessage); err != nil {
		return err
	}

	return nil
}

func (ui *UI) SetUsername(username string) {
	ui.username = username
}

func (ui *UI) SetConnection(connection *websocket.Conn) {
	ui.connection = connection
}

func (ui *UI) Connect(username string) error {
	config, err := websocket.NewConfig(WebsocketEndpoint, WebsocketOrigin)
	if err != nil {
		return err
	}

	config.Header.Set("Username", username)

	connection, err := websocket.DialConfig(config)
	if err != nil {
		return err
	}

	ui.SetConnection(connection)

	return nil
}

func (ui *UI) WriteMessage(_ *gocui.Gui, v *gocui.View) error {
	message := chat.NewMessage(chat.Regular, ui.username, v.Buffer())

	if err := websocket.JSON.Send(ui.connection, message); err != nil {
		return fmt.Errorf("UI.WriteMessage: %w", err)
	}

	v.SetCursor(0, 0)
	v.Clear()

	return nil
}

func (ui *UI) ReadMessage() error {
	for {
		var message chat.Message
		if err := websocket.JSON.Receive(ui.connection, &message); err != nil {
			return fmt.Errorf("UI.ReadMessage: %w", err)
		}

		ui.Update(func(g *gocui.Gui) error {
			switch message.Type {
			case chat.Regular:
				view, err := ui.View(MessageWidget)
				if err != nil {
					return fmt.Errorf("UI.ReadMessage: %w", err)
				}

				fmt.Fprint(view, message.Formatted())
			case chat.Connected, chat.Disconnected:
				view, err := ui.View(UsersWidget)
				if err != nil {
					return fmt.Errorf("UI.ReadMessage: %w", err)
				}

				view.Clear()
				fmt.Fprint(view, message.Text)
			}

			return nil
		})
	}

	return nil
}

func (ui *UI) Quit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

func (ui *UI) Serve() error {
	go ui.ReadMessage()

	if err := ui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return fmt.Errorf("UI.Serve: %w", err)
	}

	return nil
}
