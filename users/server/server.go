package server

import (
	"log"
	"net/http"

	"github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/proto/users"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) CreateUser(ctx context.Context, in *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	userID, err := shared.CreateUser(in.Email)

	response := new(users.CreateUserResponse)
	if err == nil {
		log.Printf("[user.Create] New user ID: %s", userID)

		createInitialItem(userID)

		// TODO: send email to user when it's created.

		response.Message = "User created successfully"
		response.Id = userID
		response.Code = http.StatusCreated
	} else {
		response.Message = err.Error()
		response.Code = http.StatusInternalServerError
	}

	return response, err
}

// Create initial item in todo list
func createInitialItem(userID string) {
	_, err := shared.ListClient.CreateItem(context.Background(), &list.CreateItemRequest{
		Message: "Welcome to Workshop!",
		UserId:  userID,
	})
	if err != nil {
		log.Printf("[user.Create] Cannot create item: %v", err)
	}
}
