package usecase

import (
	"context"
	"fmt"
	"github.com/alserov/chatter/messages/internal/db"
	"github.com/alserov/chatter/messages/internal/usecase/models"
	"time"
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
	msg.CreatedAt = time.Now()
	msg.ModifiedAt = msg.CreatedAt

	err := c.repo.CreateMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (c Chat) EditMessage(ctx context.Context, updated models.EditMessage) error {
	updated.ModifiedAt = time.Now()

	err := c.repo.EditMessage(ctx, updated)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (c Chat) DeleteMessage(ctx context.Context, deleteID string) error {
	err := c.repo.DeleteMessage(ctx, deleteID)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (c Chat) GetMessages(ctx context.Context, req models.GetParams) ([]models.Message, error) {
	msgs, err := c.repo.GetMessages(ctx, req)
	if err != nil {
		return msgs, fmt.Errorf("repo error: %w", err)
	}

	return msgs, nil
}
