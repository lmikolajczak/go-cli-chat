package data

import "time"

type Message struct {
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) SetTime(v time.Time) {
	m.Timestamp = v.Format(time.Kitchen)
}
