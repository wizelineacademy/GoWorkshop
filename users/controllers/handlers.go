package controllers

import (
	"log"

	"github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"github.com/wizelineacademy/GoWorkshop/proto/users"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"github.com/wizelineacademy/GoWorkshop/users/models"
	"golang.org/x/net/context"
)

type Service struct{}

func (s *Service) CreateUser(ctx context.Context, in *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	c := shared.DbCollection("users")
	repo := &models.UserRepository{c}

	userID, err := repo.Create(&models.User{
		Email: in.Email,
	})

	response := new(users.CreateUserResponse)
	if err == nil {
		log.Printf("[user.Create] New user ID: %s", userID)

		createInitialItem(userID)
		go notify(in.Email)

		response.Message = "User created successfully"
		response.Id = userID
		response.Code = 200
	} else {
		response.Message = err.Error()
		response.Code = 500
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

func notify(email string) {
	_, err := shared.NotifierClient.NewUser(context.Background(), &notifier.NewUserRequest{
		Email: email,
	})
	if err != nil {
		log.Printf("[user.Create] Cannot notify: %v", err)
	}
}
