package models

import "time"

type Message struct {
	ID         string
	Value      []byte
	UserID     string
	ChatID     string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Type       uint
}

const (
	TEXT  uint = iota // string
	AUDIO             // []byte
)

type EditMessage struct {
	ID         string
	Value      []byte
	ModifiedAt time.Time
}
