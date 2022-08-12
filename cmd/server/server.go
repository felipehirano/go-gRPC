package main

import (
	"net"
	"log"
	"github.com/felipehirano/go-gRPC/pb"
	"github.com/felipehirano/go-gRPC/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051");

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	// Modo reflection para rodar no evans com o seguinte comando:
	// evans -r repl --host localhost --port 50051
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}