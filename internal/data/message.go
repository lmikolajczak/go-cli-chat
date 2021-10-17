package data

import (
	"fmt"
	"time"
)

type Message struct {
	From      string `json:"from"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) SetTime(v time.Time) {
	m.Timestamp = v.Format(time.Kitchen)
}

func (m *Message) Formatted() string {
	return fmt.Sprintf("%v %v: %v", m.Timestamp, m.From, m.Text)
}
