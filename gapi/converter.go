package gapi

import (
	db "github.com/jakoo13/simplebank/db/sqlc"
	"github.com/jakoo13/simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// This file is to convert db user to pb user

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}