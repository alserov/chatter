package validate

import "github.com/MaksKazantsev/chatter/user/internal/models"

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v Validator) ValidateSignUpReq() error {
	return models.Error{}
}

func (v Validator) ValidateLoginReq() error {
	return models.Error{}
}
func (v Validator) ValidateResetPasswordReq() error {
	return models.Error{}
}
func (v Validator) ValidateSwitchNotificationsReq() error {
	return models.Error{}
}
func (v Validator) ValidateGetMainInfoReq() error {
	return models.Error{}
}

func validateLen(min, max int, s string, title string) error {
	return models.Error{}
}
