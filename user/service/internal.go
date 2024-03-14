package service

import (
	"fmt"
	"github.com/MaksKazantsev/chatter/user/internal/models"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const Salt = "jjsdjdhsh13413413"

func hashPassword(password string) (string, error) {
	fullPass := []string{password, Salt}
	in := strings.Join(fullPass, "")

	res, err := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
	if err != nil {
		return "", &models.Error{
			Msg:  fmt.Sprintf("failed to generate: %v", err),
			Code: models.ERR_INTERNAL,
		}
	}
	return string(res), nil
}

func compareAndHash(password, hash, Salt string) error {
	fullPass := []string{password, Salt}
	in := strings.Join(fullPass, "")

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(in)); err != nil {
		return &models.Error{
			Msg:  fmt.Sprintf("invalid password"),
			Code: models.ERR_CLIENT_INVALID_DATA,
		}
	}
	return nil
}
