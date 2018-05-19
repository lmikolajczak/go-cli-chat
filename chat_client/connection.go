package main

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

// Disconnect from chat and close
func Disconnect(g *gocui.Gui, v *gocui.View) error {
	connection.Close()
	return gocui.ErrQuit
}

// Send message
func Send(g *gocui.Gui, v *gocui.View) error {
	writer.WriteString(v.Buffer())
	writer.Flush()
	g.Update(func(g *gocui.Gui) error {
		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)
		return nil
	})
	return nil
}

// Connect to the server, create new reader, writer set client name
func Connect(g *gocui.Gui, v *gocui.View) error {
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	connection = conn
	reader = bufio.NewReader(connection)
	writer = bufio.NewWriter(connection)
	writer.WriteString("/name>" + v.Buffer())
	writer.Flush()
	// Some UI changes
	g.SetViewOnTop("messages")
	g.SetViewOnTop("users")
	g.SetViewOnTop("input")
	g.SetCurrentView("input")
	// Wait for server messages in new goroutine
	messagesView, _ := g.View("messages")
	usersView, _ := g.View("users")
	go func() {
		for {
			data, _ := reader.ReadString('\n')
			msg := strings.TrimSpace(data)
			switch {
			case strings.HasPrefix(msg, "/clients>"):
				data := strings.SplitAfter(msg, ">")[1]
				clientsSlice := strings.Split(data, " ")
				clientsCount := len(clientsSlice)
				var clients string
				for _, client := range clientsSlice {
					clients += client + "\n"
				}
				g.Update(func(g *gocui.Gui) error {
					usersView.Title = fmt.Sprintf(" %d users: ", clientsCount)
					usersView.Clear()
					fmt.Fprintln(usersView, clients)
					return nil
				})
			default:
				g.Update(func(g *gocui.Gui) error {
					fmt.Fprintln(messagesView, msg)
					return nil
				})
			}
		}
	}()
	return nil
}
