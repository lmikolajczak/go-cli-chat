package data

import "golang.org/x/net/websocket"

type User struct {
	Name       string          `json:"name"`
	Connection *websocket.Conn `json:"-"`
}

func NewUser(connection *websocket.Conn, name string) *User {
	return &User{
		Name:       name,
		Connection: connection,
	}
}

func (u *User) SetName(username string) {
	u.Name = username
}
