package data

type EventType int8

const (
	ConnectEvent    EventType = iota // Has the value 0.
	DisconnectEvent                  // Has the value 1.
)

type Event struct {
	Type EventType
	User *User
}

func NewEvent(_type EventType, user *User) *Event {
	return &Event{
		Type: _type,
		User: user,
	}
}
