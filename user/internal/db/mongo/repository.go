package postgres

import (
	"context"
	"github.com/MaksKazantsev/chatter/user/internal/db"
	"github.com/MaksKazantsev/chatter/user/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRepository(c *mongo.Client) db.Repository {
	return &repository{
		db: c.Database("user"),
	}
}

type repository struct {
	db *mongo.Database
}

// Friendship

func (r repository) SuggestFriendship(ctx context.Context, sender string, receiver string) error {
	panic("implement me")
}
func (r repository) RefuseFriendship(ctx context.Context, receiver string, sender string) error {
	panic("implement me")
}
func (r repository) AcceptFriendship(ctx context.Context, receiver string, sender string) error {
	panic("implement me")
}
func (r repository) GetFriendshipSuggestions(ctx context.Context, userID string) (models.UserIDs, error) {
	panic("implement me")
}

// Friends

func (r repository) GetAllFriends(ctx context.Context, userID string) (models.UserIDs, error) {
	panic("implement me")
}
func (r repository) GetFriend(ctx context.Context, userID, friendID string) (models.UserIDs, error) {
	panic("implement me")
}

// Auth

func (r repository) Signup(ctx context.Context, req models.SignupReq) error {
	panic("implement me")
}
func (r repository) Login(ctx context.Context, req models.LoginReq) (string, string, error) {
	panic("implement me")
}
func (r repository) ResetPassword(ctx context.Context, req models.ResetPasswordReq) error {
	panic("implement me")
}

// Switch

func (r repository) SwitchNotificationsStatus(ctx context.Context, userID string) error {
	panic("implement me")
}
func (r repository) SwitchPrivatePolicy(ctx context.Context, userID string) error {
	panic("implement me")
}
func (r repository) GetUserInfo(ctx context.Context, uuid string) (models.MainInfo, error) {
	panic("implement me")
}
