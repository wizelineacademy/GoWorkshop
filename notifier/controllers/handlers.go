package controllers

import (
	"log"

	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
)

type Service struct{}

func (s *Service) NewUser(ctx context.Context, in *notifier.NewUserRequest) (*notifier.NewUserResponse, error) {
	err := shared.SendEmail(in.Email, "Welcome to Wizeline Golang Workshop!")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("NewUser email sent to %s\n", in.Email)
	}

	return &notifier.NewUserResponse{
		Message: "OK",
		Code:    200,
	}, nil
}
