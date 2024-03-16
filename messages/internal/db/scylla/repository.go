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
	query := `UPDATE messages_text SET value = ?,  modified_at = ? WHERE id = ?`

	err := s.session.Query(query, []string{edit.Value, edit.ModifiedAt.String(), edit.ID}).Exec()
	if err != nil {
		// TODO: custom error
		return err
	}

	return nil
}

func (s Scylla) DeleteMessage(ctx context.Context, delete models.DeleteMessage) error {
	var query string
	switch delete.Type {
	case models.TEXT:
		query = `DELETE FROM messages_text WHERE id = ?`
	case models.AUDIO:
		query = `DELETE FROM messages_audio WHERE id = ?`
	default:
		// TODO: custom error
		return nil
	}

	err := s.session.Query(query, []string{delete.ID}).Exec()
	if err != nil {
		// TODO: custom error
		return err
	}

	return nil
}

func (s Scylla) CreateMessage(ctx context.Context, msg models.Message) error {
	var query string
	switch msg.Type {
	case models.TEXT:
		query = `INSERT INTO messages_text (id, chat_id, user_id, value, sent_at, modified_at)`
	case models.AUDIO:
		query = `INSERT INTO messages_audio (id, chat_id, user_id, value, sent_at, modified_at)`
	default:
		// TODO: custom error
		return nil
	}

	err := s.session.Query(query, []string{msg.ID, msg.ChatID, msg.UserID, msg.Value, msg.SentAt.String(), msg.ModifiedAt.String()}).Exec()
	if err != nil {
		// TODO: custom error
		return err
	}

	return nil
}
