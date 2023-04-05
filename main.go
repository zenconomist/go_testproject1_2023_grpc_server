package main

import (
	"fmt"
	ls "logstream"
	logevents "logstream1_proto"

	"net"

	"google.golang.org/grpc"
)

func main() {
	listenAddr := "localhost:50051"

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Printf("Failed to listen on %s: %v\n", listenAddr, err)
		return
	}

	grpcServer := grpc.NewServer()
	logevents.RegisterLogEventServiceServer(grpcServer, &ls.LogEventServer{})

	fmt.Printf("Starting gRPC server on %s\n", listenAddr)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC server: %v\n", err)
	}

}
