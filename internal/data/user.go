package data

import "golang.org/x/net/websocket"

type User struct {
	Name       string
	Connection *websocket.Conn
}

func NewUser(connection *websocket.Conn) *User {
	return &User{
		Connection: connection,
	}
}

func (u *User) SetName(username string) {
	u.Name = username
}
