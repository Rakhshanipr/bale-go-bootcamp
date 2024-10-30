package main

import (
	"bale-go-bootcamp/gogrpc/grpc2"
	"context"
	"log"
)

type messagingService struct {
	grpc2.MessagingServer
}

func (s *messagingService) mustEmbedUnimplementedMessagingServer() {
	log.Print("Unimplemented Messaging Server")
}

func (s *messagingService) SendMessage(ctx context.Context, req *grpc2.RequestSendMessage) (*grpc2.ResponseSendMessage, error) {

	return &grpc2.ResponseSendMessage{}, nil
}

func (s *messagingService) FetchMessage(ctx context.Context, req *grpc2.RequestFetchMessage) (*grpc2.ResponseFetchMessage, error) {
	return &grpc2.ResponseFetchMessage{}, nil
}

func (s *messagingService) GetUserMessages(ctx context.Context, req *grpc2.RequestGetUserMessages) (*grpc2.ResponseGetUserMessages, error) {
	return &grpc2.ResponseGetUserMessages{}, nil
}
