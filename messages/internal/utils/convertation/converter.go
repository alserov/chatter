package convertation

import (
	"github.com/alserov/chatter/messages/internal/usecase/models"
	messages "github.com/alserov/chatter/messages/pkg/proto/gen"
)

type Converter struct {
}

func NewConverter() Converter {
	return Converter{}
}

func (c Converter) ToMessage(in *messages.Message) models.Message {
	return models.Message{
		ID:     in.Id,
		Value:  in.Value,
		UserID: in.UserId,
		ChatID: in.ChatId,
		Type:   uint(in.Type),
	}
}

func (c Converter) ToEditMessage(in *messages.Edit) models.EditMessage {
	return models.EditMessage{
		ID:    in.Id,
		Value: []byte(in.Value),
	}
}
