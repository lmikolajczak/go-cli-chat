package data

type NotificationType int8

const (
	ConnectedUsers NotificationType = iota
)

type Notification struct {
	Type    NotificationType `json:"type"`
	Payload string           `json:"payload"`
}

func NewNotification(_type NotificationType, payload string) *Notification {
	return &Notification{
		Type:    _type,
		Payload: payload,
	}
}

func (n *Notification) Formatted() string {
	switch n.Type {
	case ConnectedUsers:
		return n.Payload
	}
	return ""
}
