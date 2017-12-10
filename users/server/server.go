package server

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/proto/notifier"
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
		go notify(in.Email)

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

// Call notifier service
func notify(email string) {
	conn, err := grpc.Dial(os.Getenv("SRV_NOTIFIER_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Error("cannot dial notifier service")
		return
	}

	_, err = notifier.NewNotifierClient(conn).Email(context.Background(), &notifier.EmailRequest{
		Email: email,
	})

	if err != nil {
		log.WithError(err).Error("unable to notifier user")
	} else {
		log.WithField("email", email).Error("user notified")
	}
}

// Create initial item in todo list
func createInitialItem(userID string) {
	conn, err := grpc.Dial(os.Getenv("SRV_LIST_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Error("cannot dial list service")
		return
	}

	if _, err = list.NewListClient(conn).CreateItem(context.Background(), &list.CreateItemRequest{
		Message: "Welcome to Workshop!",
		UserId:  userID,
	}); err != nil {
		log.WithError(err).Error("unable to create initial item")
	}
}
