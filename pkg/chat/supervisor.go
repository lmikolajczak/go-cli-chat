package chat

import (
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

type Supervisor struct {
	Users []*User
}

func NewSupervisor() *Supervisor {
	return &Supervisor{
		Users: make([]*User, 0),
	}
}

func (s *Supervisor) Join(user *User) {
	s.Users = append(s.Users, user)

	notification := NewMessage(Connected, "System", s.CurrentUsers())
	notification.SetTime(time.Now())

	s.Broadcast(notification)
}

func (s *Supervisor) Quit(user *User) {
	for i := len(s.Users) - 1; i >= 0; i-- {
		if s.Users[i] == user {
			s.Users = append(s.Users[:i], s.Users[i+1:]...)
		}
	}

	notification := NewMessage(Disconnected, "System", s.CurrentUsers())
	notification.SetTime(time.Now())

	s.Broadcast(notification)
}

func (s *Supervisor) CurrentUsers() string {
	var users string
	for _, user := range s.Users {
		users += fmt.Sprintf("%s\n", user.Name)
	}

	return users
}

func (s *Supervisor) Broadcast(message *Message) error {
	for _, user := range s.Users {
		user.Write(message)
	}

	return nil
}

func (s *Supervisor) ServeWS() func(connection *websocket.Conn) {
	return func(connection *websocket.Conn) {
		user := NewUser(connection.Request().Header.Get("Username"), connection, s)
		s.Join(user)

		user.Read()
	}
}
