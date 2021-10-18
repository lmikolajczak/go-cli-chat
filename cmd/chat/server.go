package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Luqqk/go-cli-chat/internal/data"

	"golang.org/x/net/websocket"
)

type chat struct {
	users []*data.User
	emit  chan *data.Message
	event chan *data.Event
}

func (chat *chat) serve() {
	srv := http.Server{
		Addr:    ":5000",
		Handler: chat.mux(),
	}
	go chat.run()
	srv.ListenAndServe()
}

func (chat *chat) mux() http.Handler {
	mux := http.NewServeMux()
	// Use websocket.Server because we want to accept non-browser clients,
	// which do not send an Origin header. websocket.Handler does check
	// the Origin header by default.
	mux.Handle("/", websocket.Server{
		Handler: chat.handler(),
		// Set a Server.Handshake to nil - does not check the origin.
		// We can always provide a custom handshake method to access
		// the handshake http request and implement origin check or
		// other custom logic before the connection is established.
		Handshake: nil,
	})

	return mux
}

func (chat *chat) handler() func(*websocket.Conn) {
	return func(connection *websocket.Conn) {
		user := data.NewUser(connection)
		chat.event <- data.NewEvent(data.ConnectEvent, user)

		for {
			message := data.NewMessage()
			err := websocket.JSON.Receive(user.Connection, message)
			if err != nil {
				// EOF connection closed by the client
				chat.event <- data.NewEvent(data.DisconnectEvent, user)
				return
			}
			message.SetTime(time.Now())
			chat.emit <- message
		}
	}
}

func (chat *chat) join(user *data.User) {
	chat.users = append(chat.users, user)
}

func (chat *chat) disconnect(user *data.User) {
	user.Connection.Close()
	for i := len(chat.users) - 1; i >= 0; i-- {
		if chat.users[i] == user {
			chat.users = append(chat.users[:i], chat.users[i+1:]...)
		}
	}
}

func (chat *chat) broadcast(message *data.Message) {
	for _, user := range chat.users {
		err := websocket.JSON.Send(user.Connection, message)
		if err != nil {
			log.Println(err)
		}
	}
}

func (chat *chat) run() {
	for {
		select {
		case message := <-chat.emit:
			chat.broadcast(message)
		case event := <-chat.event:
			switch event.Type {
			case data.ConnectEvent:
				chat.join(event.User)
			case data.DisconnectEvent:
				chat.disconnect(event.User)
			}
		}
	}
}
