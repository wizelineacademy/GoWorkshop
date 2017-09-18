package controllers

import (
	"log"

	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
)

type Service struct{}

func (s *Service) NewUser(ctx context.Context, in *notifier.NewUserRequest) (*notifier.NewUserResponse, error) {
	log.Printf("Sending email to %s...\n", in.Email)

	err := shared.SendEmail(in.Email, "Wizeline - Microservices with Go", "Welcome to Wizeline Go Workshop! https://github.com/wizelineacademy/GoWorkshop")

	if err != nil {
		log.Println(err.Error())
	}

	return &notifier.NewUserResponse{
		Message: "OK",
		Code:    200,
	}, nil
}
