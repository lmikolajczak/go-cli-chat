package data

type NotificationType int8

const (
	ConnectedUsers NotificationType = iota
)

type Notification struct {
	Type    NotificationType `json:"type"`
	Payload interface{}      `json:"payload"`
}

func NewNotification(_type NotificationType, payload interface{}) *Notification {
	return &Notification{
		Type:    _type,
		Payload: payload,
	}
}
