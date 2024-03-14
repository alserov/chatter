package models

type User struct {
	ID               string `bson:"id"`
	Username         string `bson:"username"`
	PhoneNumber      string `bson:"phoneNumber"`
	Photo            string `bson:"photo"`
	Password         string `bson:"password"`
	Description      string `bson:"description"`
	NotificationsOn  bool   `bson:"notificationsOn"`
	IsPrivateAccount bool   `bson:"privateAccount"`
}

type MainInfo struct {
	Username        string `bson:"username"`
	NotificationsOn bool   `bson:"notificationsOn"`
	Description     string `bson:"description"`
	PhoneNumber     string `bson:"phoneNumber"`
}

type SignupReq struct {
	ID              string `bson:"id"`
	Username        string `bson:"username"`
	PhoneNumber     string `bson:"phoneNumber"`
	Photo           string `bson:"photo"`
	Password        string `bson:"password"`
	Description     string `bson:"description"`
	NotificationsOn bool   `bson:"notificationsOn"`
}

type LoginReq struct {
	PhoneNumber string `bson:"phoneNumber"`
	Password    string `bson:"password"`
}

type ResetPasswordReq struct {
	PhoneNumber      string `bson:"phoneNumber"`
	PreviousPassword string `bson:"previousPassword"`
	NewPassword      string `bson:"newPassword"`
}

type FriendShipReq struct {
	Token  string
	UserID string `bson:"id"`
}

type UserIDs []string
type Friends struct {
	FriendIDs []string `bson:"friendIDs"`
}
