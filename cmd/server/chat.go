package main

import (
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

type message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type chat struct {
	connections []*websocket.Conn
	emit        chan message
	mutex       sync.Mutex
}

func (chat *chat) serve() {
	srv := http.Server{
		Addr:    ":5000",
		Handler: chat.mux(),
	}
	go chat.broadcast()
	srv.ListenAndServe()
}

func (chat *chat) mux() http.Handler {
	mux := http.NewServeMux()
	// Use websocket.Server because we want to accept non-browser clients,
	// which do not send an Origin header. websocket.Handler does check
	// the Origin header by default.
	mux.Handle("/", websocket.Server{
		Handler: chat.connect(),
		// Set a Server.Handshake to nil - does not check the origin.
		// We can always provide a custom handshake method to access
		// the handshake http request and implement origin check or
		// other custom logic before the connection is established.
		Handshake: nil,
	})

	return mux
}

func (chat *chat) connect() func(*websocket.Conn) {
	return func(connection *websocket.Conn) {
		chat.mutex.Lock()
		chat.connections = append(chat.connections, connection)
		chat.mutex.Unlock()

		for {
			message := message{}
			err := websocket.JSON.Receive(connection, &message)
			if err != nil {
				chat.disconnect(connection)
				return
			}
			chat.emit <- message
		}
	}
}

func (chat *chat) disconnect(connection *websocket.Conn) {
	connection.Close()
	chat.mutex.Lock()
	for i := len(chat.connections) - 1; i >= 0; i-- {
		if chat.connections[i] == connection {
			chat.connections = append(chat.connections[:i], chat.connections[i+1:]...)
		}
	}
	chat.mutex.Unlock()
}

func (chat *chat) broadcast() {
	for message := range chat.emit {
		for _, connection := range chat.connections {
			err := websocket.JSON.Send(connection, message)
			if err != nil {
				chat.disconnect(connection)
			}
		}
	}
}
