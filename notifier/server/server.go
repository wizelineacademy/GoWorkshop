package server

import (
	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) Email(ctx context.Context, in *notifier.EmailRequest) (*notifier.EmailResponse, error) {
	err := shared.SendEmail(in.Email, "Welcome to Go Workshop", "GitHub repository: https://github.com/wizelineacademy/GoWorkshop")

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
