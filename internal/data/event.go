package data

import (
	"golang.org/x/net/websocket"
)

type EventType int8

const (
	ConnectEvent    EventType = iota // Has the value 0.
	DisconnectEvent                  // Has the value 1.
)

type Event struct {
	Type       EventType
	Connection *websocket.Conn
}

func NewEvent(_type EventType, connection *websocket.Conn) *Event {
	return &Event{
		Type:       _type,
		Connection: connection,
	}
}
