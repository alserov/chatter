package models

import "time"

type Message struct {
	ID         string
	Value      []byte
	UserID     string
	ChatID     string
	SentAt     time.Time
	ModifiedAt time.Time
	Type       uint
}

const (
	TEXT  uint = iota // string
	AUDIO             // []byte
)

type EditMessage struct {
	ID         string
	Value      string
	ModifiedAt time.Time
}