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
	return models.Message{}
}

func (c Converter) ToEditMessage(in *messages.Edit) models.EditMessage {
	return models.EditMessage{}
}

func (c Converter) ToDeleteMessage(in *messages.Delete) models.DeleteMessage {
	return models.DeleteMessage{}
}
