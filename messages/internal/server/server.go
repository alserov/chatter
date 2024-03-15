package server

import (
	messages "github.com/alserov/chatter/messages/pkg/proto/gen"
	"net"
)

type Server interface {
	Serve(lis net.Listener) error
	Stop()
	messages.MessagesServer
}
