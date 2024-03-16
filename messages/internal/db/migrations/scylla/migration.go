package scylla

import (
	"github.com/scylladb/gocqlx/v2"
	"log"
)

const (
	messagesKS = `CREATE KEYSPACE IF NOT EXISTS messages
        WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1}`
	messagesList = `CREATE TABLE IF NOT EXISTS messages.list (
        ID text,
        UserID text,
        ChatID text,
        Value blob,
        SentAt timestamp,
        UpdatedAt timestamp,
        Type tinyint,
        PRIMARY KEY (ID))`
)

func MustMigrate(session gocqlx.Session) {
	if err := session.Query(messagesKS, nil).Exec(); err != nil {
		log.Fatal("Error creating keyspace:", err)
	}

	if err := session.Query(messagesList, nil).Exec(); err != nil {
		log.Fatal("Error creating table:", err)
	}
}
