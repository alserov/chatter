package convertation

import (
	"github.com/alserov/chatter/messages/internal/usecase/models"
	messages "github.com/alserov/chatter/messages/pkg/proto/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		Type:   in.Type,
	}
}

func (c Converter) ToEditMessage(in *messages.Edit) models.EditMessage {
	return models.EditMessage{
		ID:    in.Id,
		Value: []byte(in.Value),
	}
}

func (c Converter) ToGetMessages(in *messages.GetMessagesReq) models.GetParams {
	return models.GetParams{
		ChatID: in.ChatId,
		From:   in.From.AsTime(),
		To:     in.To.AsTime(),
	}
}

func (c Converter) FromMessage(in models.Message) *messages.Message {
	return &messages.Message{
		Id:         in.ID,
		Value:      in.Value,
		UserId:     in.UserID,
		ChatId:     in.ChatID,
		CreatedAt:  timestamppb.New(in.CreatedAt),
		ModifiedAt: timestamppb.New(in.ModifiedAt),
		Type:       in.Type,
	}
}

func (c Converter) FromMessages(in []models.Message) *messages.MessagesRes {
	res := &messages.MessagesRes{
		Messages: make([]*messages.Message, 0, len(in)),
	}
	for _, msg := range in {
		res.Messages = append(res.Messages, c.FromMessage(msg))
	}

	return res
}
