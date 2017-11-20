package web

import (
	"net/http"
	"os"

	"google.golang.org/grpc"

	"github.com/gocraft/web"
	"github.com/gorilla/context"

	log "github.com/sirupsen/logrus"
	pbList "github.com/wizelineacademy/GoWorkshop/proto/list"
	pbUsers "github.com/wizelineacademy/GoWorkshop/proto/users"
)

// Context struct
type Context struct {
	UsersService pbUsers.UsersClient
	ListService  pbList.ListClient
}

// ListenAndServe func
func ListenAndServe() {
	connUsers, errUsers := grpc.Dial(os.Getenv("SRV_USERS_ADDR"), grpc.WithInsecure())
	if errUsers != nil {
		log.WithError(errUsers).Fatal("could not connect to users service")
	}

	connList, errList := grpc.Dial(os.Getenv("SRV_LIST_ADDR"), grpc.WithInsecure())
	if errList != nil {
		log.WithError(errList).Fatal("could not connect to list service")
	}

	ctx := new(Context)
	ctx.UsersService = pbUsers.NewUsersClient(connUsers)
	ctx.ListService = pbList.NewListClient(connList)

	r := web.New(Context{}).
		Get("/", ctx.home).
		Post("/", ctx.createUser).
		Get("/user/:id", ctx.user).
		Post("/user/:id", ctx.user)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", r)

	log.Info("web service started")

	http.ListenAndServe(":8080", context.ClearHandler(serveMux))
}
