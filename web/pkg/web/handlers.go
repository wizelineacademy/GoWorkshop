package web

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/gocraft/web"
	"github.com/gorilla/context"

	pbUsers "github.com/wizelineacademy/GoWorkshop/proto/users"
)

// Context struct
type Context struct {
	UsersService pbUsers.UsersClient
}

// ListenAndServe func
func ListenAndServe() {
	conn, err := grpc.Dial("users:8080", grpc.WithInsecure())
	if err != nil {
		log.Printf("cannot connect to users service: %v", err)
	}

	ctx := new(Context)
	ctx.UsersService = pbUsers.NewUsersClient(conn)

	r := web.New(Context{}).
		Get("/", ctx.home).
		Get("/user", ctx.home).
		Post("/user", ctx.user)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", r)
	staticFolders := []string{"scripts", "styles"}
	for _, sf := range staticFolders {
		serveMux.Handle("/"+sf+"/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	}

	http.ListenAndServe(":8080", context.ClearHandler(serveMux))
}

// Ajax func
func (c *Context) Ajax(w web.ResponseWriter, r *web.Request, response interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resultJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJSON)
}
