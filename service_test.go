package main

import (
	"context"
	"crypto/md5"
	"testing"

	"github.com/realbucksavage/simple-user-service/generated/users"
	"github.com/stretchr/testify/require"
)

func TestDefaultUserService_GetUser(t *testing.T) {

	svc := &defaultUserService{
		usersMap: map[string]*users.User{
			"user001": {
				UserID:   "user001",
				Name:     "Test",
				Email:    "Test",
				Password: []byte("APassword"),
			},
		},
	}

	pwd := md5.Sum([]byte("APassword"))

	ctx := context.Background()

	u, err := svc.GetUser(ctx, &users.GetUserRequest{UserID: "user001"})
	require.NoError(t, err)
	require.Equal(t, u.Password, pwd[:])

	u, err = svc.GetUser(ctx, &users.GetUserRequest{UserID: "Bogus"})
	require.Error(t, err)
	require.Nil(t, u)
}
