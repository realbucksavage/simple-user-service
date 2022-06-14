package main

import (
	"context"
	"crypto/md5"
	"log"

	"github.com/realbucksavage/simple-user-service/generated/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var userRegistry = map[string]*users.User{
	"user0001": {
		UserID:   "user0001",
		Name:     "User 1",
		Email:    "user1@acme.com",
		Password: []byte("SomePassword1"),
	},
	"user0002": {
		UserID:   "user0002",
		Name:     "User 2",
		Email:    "user2@acme.com",
		Password: []byte("SomePassword2"),
	},
	"user0003": {
		UserID:   "user0003",
		Name:     "User 3",
		Email:    "user3@acme.com",
		Password: []byte("SomePassword3"),
	},
}

type defaultUserService struct {
	users.UnimplementedUserServiceServer
	usersMap map[string]*users.User
}

func (d *defaultUserService) GetUser(_ context.Context, req *users.GetUserRequest) (*users.User, error) {

	u, ok := d.usersMap[req.UserID]
	if !ok {
		log.Printf("user with ID %q not found", req.UserID)
		return nil, status.Errorf(codes.NotFound, "user %q for found", req.UserID)
	}

	sum := md5.Sum(u.Password)
	u.Password = sum[:]

	return u, nil
}
