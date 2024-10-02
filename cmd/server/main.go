package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/parampreetr/grpc_go/api/gen"
	"github.com/parampreetr/grpc_go/config"
	"github.com/parampreetr/grpc_go/db"
	"github.com/parampreetr/grpc_go/internal/services"
	"google.golang.org/grpc"
)

func main() {
	config.LoadEnvVariables()
	env := config.GetEnvConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", env.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	DB, err := db.InitDB()

	if err != nil {
		log.Fatalf("unable to initilize Database: %v", err)
	}

	pb.RegisterTaskServiceServer(grpcServer, &services.GRPCServer{DB: DB})
	log.Printf("Started GPRC Server on port %d\n", env.ServerPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
