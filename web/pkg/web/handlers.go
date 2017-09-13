package web

import (
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/gocraft/web"
	"github.com/gorilla/context"

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
	connUsers, errUsers := grpc.Dial("users:8080", grpc.WithInsecure())
	if errUsers != nil {
		log.Fatalf("cannot connect to users service: %v", errUsers)
	}

	connList, errList := grpc.Dial("list:8080", grpc.WithInsecure())
	if errList != nil {
		log.Fatalf("[web] cannot connect to list service: %v", errList)
	}

	ctx := new(Context)
	ctx.UsersService = pbUsers.NewUsersClient(connUsers)
	ctx.ListService = pbList.NewListClient(connList)

	r := web.New(Context{}).
		Get("/", ctx.home).
		Post("/", ctx.createUser).
		Get("/user/:id", ctx.user)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", r)

	http.ListenAndServe(":8080", context.ClearHandler(serveMux))
}
