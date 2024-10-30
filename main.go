package main

import (
	grpc2 "bale-go-bootcamp/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:4041")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service := messagingService{}
	grpc2.RegisterMessagingServer(server, &service)
	server.Serve(lis)

}
