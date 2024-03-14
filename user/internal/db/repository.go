package db

import (
	"context"
	"github.com/MaksKazantsev/chatter/user/internal/models"
)

type Repository interface {
	SuggestFriendship(ctx context.Context, sender string, receiver string) error
	RefuseFriendship(ctx context.Context, receiver string, sender string) error
	AcceptFriendship(ctx context.Context, receiver string, sender string) error
	GetAllFriends(ctx context.Context, userID string) (models.UserIDs, error)
	GetFriend(ctx context.Context, userID, friendID string) (models.User, error)
	GetFriendshipSuggestions(ctx context.Context, userID string) (models.UserIDs, error)

	Signup(ctx context.Context, req models.SignupReq) error
	Login(ctx context.Context, req models.LoginReq) (string, string, error)
	ResetPassword(ctx context.Context, req models.ResetPasswordReq) error

	SwitchNotificationsStatus(ctx context.Context, userID string) error
	SwitchPrivatePolicy(ctx context.Context, userID string) error
	GetUserInfo(ctx context.Context, userID string) (models.MainInfo, error)
}
