package webservers

import (
	"calendar_service/application/configs"
	"calendar_service/frameworks/logger"
	"context"
	"net"

	grpc_impl "calendar_service/application/services/api/grpc"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Log incoming gRPC requests
	logger.Info("Received gRPC request", zap.String("method", info.FullMethod))

	// Call the handler to process the request
	resp, err := handler(ctx, req)

	// Log outgoing gRPC responses
	logger.Info("Sending gRPC response", zap.String("method", info.FullMethod))

	return resp, err
}

func Serve(config *configs.ServerConfig) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatal("failed to listen:", err)
	}

	grpc_impl.Config(&config.Storage)

	// Create a new gRPC server
	srv := grpc.NewServer(
		// Register the logging interceptor
		grpc.UnaryInterceptor(loggingInterceptor),
	)

	// Enable reflection on the server
	reflection.Register(srv)
	grpc_impl.InitEventServer(srv)
	if err := srv.Serve(lis); err != nil {
		logger.Fatal("failed to serve:", err)
	}
}
