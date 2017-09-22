package handlers

import (
	"log"

	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
)

type Service struct{}

func (s *Service) Email(ctx context.Context, in *notifier.EmailRequest) (*notifier.EmailResponse, error) {
	log.Printf("Sending email to %s...\n", in.Email)

	err := shared.SendEmail(in.Email, "Wizeline - Microservices with Go", "Welcome to Wizeline Go Workshop! https://github.com/wizelineacademy/GoWorkshop")

	if err != nil {
		log.Println(err.Error())
	}

	return &notifier.EmailResponse{
		Message: "OK",
		Code:    200,
	}, nil
}
