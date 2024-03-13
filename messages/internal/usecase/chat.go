package usecase

import (
	"context"
	"fmt"
	"github.com/alserov/chatter/messages/internal/repository"
	"github.com/alserov/chatter/messages/internal/usecase/models"
)

func NewChat() *Chat {
	return &Chat{}
}

type Chat struct {
	repo repository.Repository
}

func (c Chat) CreateMessage(ctx context.Context, msg models.Message) error {
	err := c.repo.CreateMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (c Chat) EditMessage(ctx context.Context, updated models.EditMessage) error {
	err := c.repo.EditMessage(ctx, updated)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (c Chat) DeleteMessage(ctx context.Context, delete models.DeleteMessage) error {
	err := c.repo.DeleteMessage(ctx, delete)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}
