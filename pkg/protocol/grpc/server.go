package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"
	v1 "todo/pkg/api/v1"

	"todo/pkg/logger"

	"google.golang.org/grpc"
)

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logger.Log.Warn("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()
	// start gRPC server
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}
