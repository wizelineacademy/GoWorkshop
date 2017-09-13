package web

import (
	"context"

	"github.com/wizelineacademy/GoWorkshop/web/pkg/tpl"

	"github.com/gocraft/web"
	pbUsers "github.com/wizelineacademy/GoWorkshop/proto/users"
)

func (c *Context) home(w web.ResponseWriter, r *web.Request) {
	d := tpl.Data{
		TemplateFile: "pages/home.html",
	}

	d.Render(w, r)
}

func (c *Context) user(w web.ResponseWriter, r *web.Request) {
	email := r.FormValue("email")
	name := r.FormValue("name")

	// gRPC call
	resp, err := c.UsersService.CreateUser(context.Background(), &pbUsers.CreateUserRequest{
		Email: email,
		Name:  name,
	})
	if err != nil {
		resp = &pbUsers.CreateUserResponse{
			Message: err.Error(),
		}
	}

	d := tpl.Data{
		TemplateFile: "pages/user.html",
		Data: struct {
			ID       string
			Response *pbUsers.CreateUserResponse
		}{
			ID:       resp.GetId(),
			Response: resp,
		},
	}

	d.Render(w, r)
}
