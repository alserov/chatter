package models

type StatusCode uint

const (
	ERR_INTERNAL StatusCode = iota
	ERR_CLIENT_INVALID_DATA
	ERR_CLIENT_NOT_FOUND
	ERR_NOT_ALLOWED
)

type Error struct {
	Msg  string
	Code StatusCode
}

func (e Error) Error() string {
	return e.Msg
}
