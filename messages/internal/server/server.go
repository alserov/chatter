package server

import (
	"context"
	"github.com/alserov/chatter/messages/internal/usecase"
	"github.com/alserov/chatter/messages/internal/utils/convertation"
	"github.com/alserov/chatter/messages/internal/utils/validation"
	messages "github.com/alserov/chatter/messages/pkg/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type Server interface {
	Serve(lis net.Listener) error
	Stop()
	messages.MessagesServer
}

func NewServer(ucase usecase.Chat) Server {
	grpcServer := grpc.NewServer()
	server := &server{
		srvr: grpcServer,
		uc:   ucase,
	}

	messages.RegisterMessagesServer(grpcServer, server)

	return server
}

type server struct {
	srvr *grpc.Server

	uc       usecase.Chat
	convert  convertation.Converter
	validate validation.Validator

	messages.UnimplementedMessagesServer
}

func (s server) Serve(lis net.Listener) error {
	return s.srvr.Serve(lis)
}

func (s server) Stop() {
	s.srvr.GracefulStop()
}

func (s server) CreateMessage(ctx context.Context, message *messages.Message) (*emptypb.Empty, error) {
	if err := s.validate.ValidateMessage(message); err != nil {
		return nil, err
	}

	err := s.uc.CreateMessage(ctx, s.convert.ToMessage(message))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s server) DeleteMessage(ctx context.Context, delete *messages.Delete) (*emptypb.Empty, error) {
	if err := s.validate.ValidateDeleteMessage(delete); err != nil {
		return nil, err
	}

	err := s.uc.DeleteMessage(ctx, s.convert.ToDeleteMessage(delete))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s server) EditMessage(ctx context.Context, edit *messages.Edit) (*emptypb.Empty, error) {
	if err := s.validate.ValidateEditMessage(edit); err != nil {
		return nil, err
	}

	err := s.uc.EditMessage(ctx, s.convert.ToEditMessage(edit))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
