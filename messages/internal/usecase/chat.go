package usecase

import (
	"context"
	"github.com/alserov/chatter/messages/internal/usecase/models"
)

type Chat interface {
	CreateMessage(ctx context.Context, msg models.Message) error
	EditMessage(ctx context.Context, updated models.EditMessage) error
	DeleteMessage(ctx context.Context, messageID, chatID string) error
}

func NewChat() Chat {
	return &chat{}
}

type chat struct {
}

func (c chat) CreateMessage(ctx context.Context, msg models.Message) error {
	//TODO implement me
	panic("implement me")
}

func (c chat) EditMessage(ctx context.Context, updated models.EditMessage) error {
	//TODO implement me
	panic("implement me")
}

func (c chat) DeleteMessage(ctx context.Context, messageID, chatID string) error {
	//TODO implement me
	panic("implement me")
}
