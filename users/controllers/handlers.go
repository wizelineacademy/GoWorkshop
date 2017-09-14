package controllers

import (
	"log"

	pbList "github.com/wizelineacademy/GoWorkshop/proto/list"
	pb "github.com/wizelineacademy/GoWorkshop/proto/users"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"github.com/wizelineacademy/GoWorkshop/users/models"
	"golang.org/x/net/context"
)

// Service struct
type Service struct{}

// CreateUser implementation
func (s *Service) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (response *pb.CreateUserResponse, err error) {
	user := &models.User{
		Email: in.Email,
	}

	appContext := shared.NewContext()
	defer appContext.Close()

	c := appContext.DbCollection("users")
	repo := &models.UserRepository{
		C: c,
	}
	var userID string
	userID, err = repo.Create(user)

	response = new(pb.CreateUserResponse)
	if err == nil {
		log.Printf("[user.Create] New user ID: %s", userID)

		createInitialItem(appContext, userID)

		response.Message = "User created successfully"
		response.Id = userID
		response.Code = 200
	} else {
		response.Message = err.Error()
		response.Code = 500
	}

	return
}

// Create initial item in todo list
func createInitialItem(appContext *shared.Context, userID string) {
	_, listErr := appContext.ListService.CreateItem(context.Background(), &pbList.CreateItemRequest{
		Message: "Welcome to Workshop!",
		UserId:  userID,
	})
	if listErr != nil {
		log.Printf("[user.Create] Cannot create item: %v", listErr)
	}
}
