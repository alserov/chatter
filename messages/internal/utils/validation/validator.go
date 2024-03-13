package validation

import messages "github.com/alserov/chatter/messages/pkg/proto/gen"

type Validator struct {
}

func NewValidator() Validator {
	return Validator{}
}

func (v Validator) ValidateMessage(in *messages.Message) error {
	return nil
}

func (v Validator) ValidateDeleteMessage(in *messages.Delete) error {
	return nil
}

func (v Validator) ValidateEditMessage(in *messages.Edit) error {

	return nil
}
