package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	authgrpc "sso/internal/grpc/auth"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(log, gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Run() error {

	log := a.log.With(
		slog.String("op", "grpcapp.Run"),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		a.log.Warn("failed to listen tcp", slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", "grpcapp.Run", err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		a.log.Warn("failed to serve", slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", "grpcapp.Run", err)
	}

	return nil
}

func (a *App) Stop() {
	a.log.With(slog.String("op", "grpcapp.Stop")).
		Info("Stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()

}
