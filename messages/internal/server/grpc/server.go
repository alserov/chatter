package grpc

import (
	"context"
	"github.com/alserov/chatter/messages/internal/server"
	"github.com/alserov/chatter/messages/internal/usecase"
	"github.com/alserov/chatter/messages/internal/utils/convertation"
	"github.com/alserov/chatter/messages/internal/utils/validation"
	messages "github.com/alserov/chatter/messages/pkg/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

func NewServer(ucase *usecase.Chat) server.Server {
	gRPCServer := grpc.NewServer()
	s := &grpcServer{
		srvr: gRPCServer,
		uc:   ucase,
	}

	messages.RegisterMessagesServer(gRPCServer, s)

	return s
}

type grpcServer struct {
	srvr *grpc.Server

	uc       *usecase.Chat
	convert  convertation.Converter
	validate validation.Validator

	messages.UnimplementedMessagesServer
}

func (s grpcServer) Serve(lis net.Listener) error {
	return s.srvr.Serve(lis)
}

func (s grpcServer) Stop() {
	s.srvr.GracefulStop()
}

func (s grpcServer) CreateMessage(ctx context.Context, message *messages.Message) (*emptypb.Empty, error) {
	if err := s.validate.ValidateMessage(message); err != nil {
		return nil, err
	}

	err := s.uc.CreateMessage(ctx, s.convert.ToMessage(message))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s grpcServer) DeleteMessage(ctx context.Context, delete *messages.Delete) (*emptypb.Empty, error) {
	if err := s.validate.ValidateDeleteMessage(delete); err != nil {
		return nil, err
	}

	err := s.uc.DeleteMessage(ctx, delete.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s grpcServer) EditMessage(ctx context.Context, edit *messages.Edit) (*emptypb.Empty, error) {
	if err := s.validate.ValidateEditMessage(edit); err != nil {
		return nil, err
	}

	err := s.uc.EditMessage(ctx, s.convert.ToEditMessage(edit))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
