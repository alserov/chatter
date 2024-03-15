package usecase

import (
	"context"
	"fmt"
	"github.com/alserov/chatter/messages/internal/db"
	"github.com/alserov/chatter/messages/internal/usecase/models"
)

type Param struct {
	Repo db.Repository
}

func NewChat(p Param) *Chat {
	return &Chat{
		repo: p.Repo,
	}
}

type Chat struct {
	repo db.Repository
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
