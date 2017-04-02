package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
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
			client.incoming <- fmt.Sprintf("- %s disconnected\n", client.name)
			client.status = 0
			client.disconnect <- true
			client.conn.Close()
			break
		}
		switch {
		case strings.HasPrefix(msg, "/name>"):
			name := strings.TrimSpace(strings.SplitAfter(msg, ">")[1])
			client.name = name
			client.incoming <- fmt.Sprintf("+ %s connected\n", name)
		default:
			client.incoming <- fmt.Sprintf("%s: %s", client.name, msg)
		}
	}
}

// Chat struct
type Chat struct {
	clients  []*Client
	connect  chan net.Conn
	outgoing chan string
}

// CreateChat creates new chat and
// starts listening for connections.
func CreateChat() *Chat {
	chat := &Chat{
		clients:  make([]*Client, 0),
		connect:  make(chan net.Conn),
		outgoing: make(chan string),
	}

	chat.Listen()

	return chat
}

// Listen listens for connections and
// messages to broadcast.
func (chat *Chat) Listen() {
	go func() {
		for {
			select {
			case conn := <-chat.connect:
				chat.Join(conn)
			case msg := <-chat.outgoing:
				chat.Broadcast(msg)
			}
		}
	}()
}

// Connect passing connection to the chat.
func (chat *Chat) Connect(conn net.Conn) {
	chat.connect <- conn
}

// Join creates new client and starts listening
// for client messages.
func (chat *Chat) Join(conn net.Conn) {
	client := CreateClient(conn)
	chat.clients = append(chat.clients, client)
	go func() {
		for {
			chat.outgoing <- <-client.incoming
		}
	}()
}

// Broadcast sends message to all connected clients.
func (chat *Chat) Broadcast(data string) {
	currentTime := time.Now().Format("15:04:05")
	msg := fmt.Sprintf("[%s] %s", currentTime, data)
	for i, client := range chat.clients {
		if client.status == 0 {
			chat.clients = append(chat.clients[:i], chat.clients[i+1:]...)
		}
	}
	for _, client := range chat.clients {
		client.outgoing <- msg
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}

	chat := CreateChat()

	for { // listen for connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		chat.Connect(conn)
	}
}
