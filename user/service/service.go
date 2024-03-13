package service

type UserService interface {
	FriendShip
	Auth
	Switch
}

type FriendShip interface {
}
type Auth interface {
}
type Switch interface {
}
