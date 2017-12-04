package server

import (
	"github.com/wizelineacademy/GoWorkshop/notifier/smtp2go"
	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) Email(ctx context.Context, in *notifier.EmailRequest) (*notifier.EmailResponse, error) {
	err := smtp2go.SendEmail(in.Email, "Welcome to Go Workshop", "GitHub repository: https://github.com/wizelineacademy/GoWorkshop")

	if err != nil {
		return &notifier.EmailResponse{
			Message: err.Error(),
			Code:    500,
		}, nil
	}

	return &notifier.EmailResponse{
		Message: "OK",
		Code:    200,
	}, nil
}
