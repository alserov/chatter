package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/MaksKazantsev/chatter/user/internal/db"
	"github.com/MaksKazantsev/chatter/user/internal/models"
	"go.mongodb.org/mongo-driver/bson"
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
	coll := r.db.Collection("friends")

	var fr models.Friends

	filter := bson.M{
		"userID": userID,
	}

	err := coll.FindOne(ctx, filter).Decode(&fr)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil, &models.Error{
				Msg:  "no friends been found",
				Code: models.ERR_CLIENT_NOT_FOUND,
			}
		}
		return nil, &models.Error{
			Msg:  fmt.Sprintf("failed to found friends: %v", err),
			Code: models.ERR_INTERNAL,
		}
	}
	return fr.FriendIDs, nil
}

func (r repository) GetFriend(ctx context.Context, userID, friendID string) (models.User, error) {
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
	coll := r.db.Collection("users")

	filter := bson.M{
		"phoneNumber": req.PhoneNumber,
	}

	updated := bson.D{
		{Key: "$set", Value: bson.M{"password": req.NewPassword}},
	}

	_, err := coll.UpdateOne(ctx, filter, updated)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return &models.Error{
				Msg:  fmt.Sprintf("user with phone number '%s' not found", req.PhoneNumber),
				Code: models.ERR_CLIENT_NOT_FOUND,
			}
		}
		return &models.Error{
			Msg:  fmt.Sprintf("failed to update: %v", err),
			Code: models.ERR_INTERNAL,
		}
	}

	return nil
}

// Switch

func (r repository) SwitchNotificationsStatus(ctx context.Context, userID string) error {
	coll := r.db.Collection("users")

	filter := bson.M{
		"id": userID,
	}
	update := bson.M{
		"$set": bson.M{"$not": "notificationsOn"},
	}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return &models.Error{
				Msg:  "user not found",
				Code: models.ERR_CLIENT_NOT_FOUND,
			}
		}
		return &models.Error{
			Msg:  fmt.Sprintf("failed to update notifications status: %v", err),
			Code: models.ERR_INTERNAL,
		}
	}
	return nil
}

func (r repository) SwitchPrivatePolicy(ctx context.Context, userID string) error {
	coll := r.db.Collection("users")

	filter := bson.M{
		"id": userID,
	}
	update := bson.M{
		"$set": bson.M{"$not": "privateAccount"},
	}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return &models.Error{
				Msg:  "user not found",
				Code: models.ERR_CLIENT_NOT_FOUND,
			}
		}
		return &models.Error{
			Msg:  fmt.Sprintf("failed to update private policy: %v", err),
			Code: models.ERR_INTERNAL,
		}
	}
	return nil
}
func (r repository) GetUserInfo(ctx context.Context, userID string) (models.MainInfo, error) {
	coll := r.db.Collection("users")

	filter := bson.M{
		"id": userID,
	}

	var info models.MainInfo
	err := coll.FindOne(ctx, filter).Decode(&info)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return models.MainInfo{}, &models.Error{
				Msg:  "user not found",
				Code: models.ERR_CLIENT_NOT_FOUND,
			}
		}
		return models.MainInfo{}, &models.Error{
			Msg:  fmt.Sprintf("failed to update: %v", err),
			Code: models.ERR_INTERNAL,
		}
	}

	return info, nil
}
