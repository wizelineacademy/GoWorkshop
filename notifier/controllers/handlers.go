package controllers

import (
	pb "github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
)

// Service struct
type Service struct{}

// NewUser implementation
func (s *Service) NewUser(ctx context.Context, in *pb.NewUserRequest) (response *pb.NewUserResponse, err error) {
	email := in.Email
	shared.SendEmail(email, "Welcome to Wizeline Golang Workshop!")

	return &pb.NewUserResponse{
		Message: "OK",
		Code:    200,
	}, nil
}
