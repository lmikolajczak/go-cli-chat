package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Client struct
type Client struct {
	name       string
	conn       net.Conn
	writer     *bufio.Writer
	reader     *bufio.Reader
	incoming   chan string
	outgoing   chan string
	disconnect chan bool
	status     int // 1 connected, 0 otherwise
}

// CreateClient creates new client and starts listening
// for incoming and outgoing messages.
func CreateClient(conn net.Conn) *Client {
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	client := &Client{
		name:       "user",
		conn:       conn,
		writer:     writer,
		outgoing:   make(chan string),
		reader:     reader,
		incoming:   make(chan string),
		disconnect: make(chan bool),
		status:     1,
	}

	go client.Write()
	go client.Read()

	return client
}

// Write writes message to the client.
func (client *Client) Write() {
	for {
		select {
		case <-client.disconnect:
			client.status = 0
			break
		default:
			msg := <-client.outgoing
			client.writer.WriteString(msg)
			client.writer.Flush()
		}
	}
}

// Read reads message from the client.
func (client *Client) Read() {
	for {
		msg, err := client.reader.ReadString('\n')
		if err != nil {
			client.incoming <- fmt.Sprintf("\x1b[0;31m- %s disconnected\033[0m\n", client.name)
			client.status = 0
			client.disconnect <- true
			client.conn.Close()
			break
		}
		switch {
		case strings.HasPrefix(msg, "/name>"):
			name := strings.TrimSpace(strings.SplitAfter(msg, ">")[1])
			client.name = name
			client.incoming <- fmt.Sprintf("\x1b[0;32m+ %s connected\033[0m\n", name)
		default:
			client.incoming <- fmt.Sprintf("%s: %s", client.name, msg)
		}
	}
}
