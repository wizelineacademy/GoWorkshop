package controllers

import (
	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
	"log"
)

type Service struct{}

func (s *Service) NewUser(ctx context.Context, in *notifier.NewUserRequest) (*notifier.NewUserResponse, error) {
	err := shared.SendEmail(in.Email, "Welcome to Wizeline Golang Workshop!")
	if err != nil {
		log.Println(err.Error())
	}

	return &notifier.NewUserResponse{
		Message: "OK",
		Code:    200,
	}, nil
}
