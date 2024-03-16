package scylla

import (
	"context"
	"github.com/alserov/chatter/messages/internal/db"
	"github.com/alserov/chatter/messages/internal/usecase/models"
	"github.com/scylladb/gocqlx/v2"
)

var _ db.Repository = &Scylla{}

func NewRepository(s gocqlx.Session) *Scylla {
	return &Scylla{
		session: s,
	}
}

type Scylla struct {
	session gocqlx.Session
}

func (s Scylla) EditMessage(ctx context.Context, edit models.EditMessage) error {
	query := `UPDATE text SET value = ?,  modified_at = ? WHERE id = ?`

	err := s.session.Query(query, []string{string(edit.Value), edit.ModifiedAt.String(), edit.ID}).Exec()
	if err != nil {
		// TODO: custom error
		return err
	}

	return nil
}

func (s Scylla) DeleteMessage(ctx context.Context, deleteID string) error {
	query := `DELETE FROM list WHERE id = ?`

	err := s.session.Query(query, []string{deleteID}).Exec()
	if err != nil {
		// TODO: custom error
		return err
	}

	return nil
}

func (s Scylla) CreateMessage(ctx context.Context, msg models.Message) error {
	query := `INSERT INTO text (id, chat_id, user_id, value, created_at, modified_at) VALUES (?, ?, ?, ?, ?, ?)`
	err := s.session.Query(query, []string{msg.ID, msg.ChatID, msg.UserID, string(msg.Value), msg.CreatedAt.String(), msg.ModifiedAt.String()}).Exec()
	if err != nil {
		// TODO: custom error
		return err
	}

	return nil
}
