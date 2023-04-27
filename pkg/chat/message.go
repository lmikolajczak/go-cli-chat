package chat

import (
	"fmt"
	"time"
)

type MessageType int8

const (
	Regular      MessageType = iota // Has the value 0.
	Connected                       // Has the value 1.
	Disconnected                    // Has the value 2.
)

type Message struct {
	Type      MessageType `json:"type"`
	From      string      `json:"from"`
	Text      string      `json:"text"`
	Timestamp string      `json:"timestamp"`
}

func NewMessage(msgType MessageType, from string, text string) *Message {
	return &Message{
		Type:      msgType,
		From:      from,
		Text:      text,
		Timestamp: "",
	}
}

func (m *Message) SetTime(v time.Time) {
	m.Timestamp = v.Format(time.Kitchen)
}

func (m *Message) Formatted() string {
	return fmt.Sprintf("%v %v: %v", m.Timestamp, m.From, m.Text)
}
