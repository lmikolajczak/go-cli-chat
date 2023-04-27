package chat

import (
	"time"

	"golang.org/x/net/websocket"
)

type User struct {
	Name       string          `json:"name"`
	Connection *websocket.Conn `json:"-"`
	Egress     chan *Message   `json:"-"`
	Supervisor *Supervisor     `json:"-"`
}

func NewUser(name string, connection *websocket.Conn, supervisor *Supervisor) *User {
	return &User{
		Name:       name,
		Connection: connection,
		Egress:     make(chan *Message),
		Supervisor: supervisor,
	}
}

func (u *User) Read() {
	for {
		message := &Message{}
		if err := websocket.JSON.Receive(u.Connection, message); err != nil {
			// EOF connection closed by the client
			u.Supervisor.Quit(u)
			break
		}

		message.SetTime(time.Now())
		u.Supervisor.Broadcast(message)
	}
}

func (u *User) Write(message *Message) {
	if err := websocket.JSON.Send(u.Connection, message); err != nil {
		// EOF connection closed by the client
		u.Supervisor.Quit(u)
	}
}
