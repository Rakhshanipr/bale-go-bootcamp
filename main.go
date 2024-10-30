package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:4042")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("start", "start")

	server := grpc.NewServer()
	service := messagingService{}
	grpc.RegisterMessagingServer(server, &service)
	server.Serve(lis)
	defer server.GracefulStop()
}
