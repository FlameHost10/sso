package auth

import (
	"context"
	ssov1 "github.com/FlameHost10/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

const (
	emptyValue = 0
)

type Auth interface {
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	log *slog.Logger
}

func Register(log *slog.Logger, gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{
		log: log,
	})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if req.GetEmail() == "" {
		s.log.Warn("email is empty")
		return nil, status.Error(codes.InvalidArgument, "email is empty")
	}

	if req.GetPassword() == "" {
		s.log.Warn("password is empty")
		return nil, status.Error(codes.InvalidArgument, "password is empty")
	}

	if req.GetAppId() == emptyValue {
		s.log.Warn("app_id is required")
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	return &ssov1.LoginResponse{
		Token: "token123-" + req.GetEmail(),
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
