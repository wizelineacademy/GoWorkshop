package server

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/proto/users"
	"github.com/wizelineacademy/GoWorkshop/users/models"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) CreateUser(ctx context.Context, in *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	userID, err := models.CreateUser(in.Email)

	response := new(users.CreateUserResponse)
	if err == nil {
		log.WithField("id", userID).Info("user created")

		createInitialItem(userID)

		// TODO: send email to user when it's created.

		response.Message = "User created successfully"
		response.Id = userID
		response.Code = http.StatusCreated
	} else {
		log.WithError(err).Error("unable to create user")

		response.Message = err.Error()
		response.Code = http.StatusInternalServerError
	}

	return response, err
}

// Create initial item in todo list using list service
func createInitialItem(userID string) {
	conn, err := grpc.Dial(os.Getenv("SRV_LIST_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Error("cannot dial list service")
		return
	}

	_, err = list.NewListClient(conn).CreateItem(context.Background(), &list.CreateItemRequest{
		Message: "Welcome to Workshop!",
		UserId:  userID,
	})

	if err != nil {
		log.WithError(err).Error("unable to create initial item")
	}
}
