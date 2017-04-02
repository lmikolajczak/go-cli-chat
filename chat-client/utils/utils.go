package utils

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/jroimartin/gocui"
)

var (
	connection net.Conn
	reader     *bufio.Reader
	writer     *bufio.Writer
)

func Disconnect(g *gocui.Gui, v *gocui.View) error {
	connection.Close()
	return gocui.ErrQuit
}

func Send(g *gocui.Gui, v *gocui.View) error {
	writer.WriteString(v.Buffer())
	writer.Flush()
	g.Execute(func(g *gocui.Gui) error {
		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)
		return nil
	})
	return nil
}

func Connect(g *gocui.Gui, v *gocui.View) error {
	// Connect to the server
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	connection = conn
	// Create server reader and writer and set nick
	reader = bufio.NewReader(connection)
	writer = bufio.NewWriter(connection)
	writer.WriteString("/nick>" + v.Buffer())
	writer.Flush()
	// Some UI changes
	g.SetViewOnTop("messages")
	g.SetViewOnTop("input")
	g.SetCurrentView("input")
	// Wait for server messages
	messagesView, _ := g.View("messages")
	go func() {
		for {
			data, _ := reader.ReadString('\n')
			msg := strings.TrimSpace(data)
			g.Execute(func(g *gocui.Gui) error {
				fmt.Fprintln(messagesView, msg)
				return nil
			})
		}
	}()
	return nil
}
