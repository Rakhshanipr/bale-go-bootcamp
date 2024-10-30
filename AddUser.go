package main

import (
	"bale-go-bootcamp/gogrpc/grpc2"
	"context"
	"errors"
	"log"
	"unicode"
)

var users []user

func containsDigit(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

var last_user_id = 0

func containsLetter(s string) bool {
	for _, char := range s {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

func isExistUserName(username string) bool {
	for _, user := range users {
		if user.username == username {
			return true
		}
	}
	return false
}

func (s *messagingService) AddUser(ctx context.Context, req *grpc2.RequestAddUser) (*grpc2.ResponseAddUser, error) {
	log.Fatal("send message: " + req.UserName)
	if len(req.UserName) < 4 || !containsDigit(req.UserName) || !containsLetter(req.UserName) || isExistUserName(req.UserName) {
		return nil, errors.New("invalid username")
	}

	return &grpc2.ResponseAddUser{
		Id: 9,
	}, nil
	//last_user_id++
	//users = append(users, user{
	//	id:       last_user_id,
	//	username: req.UserName,
	//	file_id:  req.FileId,
	//})
	//
	//return &grpc2.ResponseAddUser{
	//	Id: int32(last_user_id),
	//}, nil
}
