package repository

import (
	"context"
	"github.com/alserov/chatter/messages/internal/usecase/models"
)

type Repository interface {
	CreateMessage(ctx context.Context, msg models.Message) error
	EditMessage(ctx context.Context, edit models.EditMessage) error
	DeleteMessage(ctx context.Context, delete models.DeleteMessage) error
}
