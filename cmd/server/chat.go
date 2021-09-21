package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type chat struct{}

func (chat *chat) serve() {
	srv := http.Server{
		Addr:    ":5000",
		Handler: chat.mux(),
	}
	srv.ListenAndServe()
}

func (chat *chat) mux() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", websocket.Server{
		Handler: chat.connect(),
	})

	return mux
}

func (chat *chat) connect() func(*websocket.Conn) {
	return func(connection *websocket.Conn) {
		message := message{}
		for {
			err := websocket.JSON.Receive(connection, &message)
			if err != nil {
				connection.Close()
				return
			}
			log.Println("new message:", message.Type, message.Text, message)
		}
	}
}
