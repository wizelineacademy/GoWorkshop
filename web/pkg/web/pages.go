package web

import (
	"context"
	"net/http"

	"github.com/wizelineacademy/GoWorkshop/web/pkg/tpl"

	"github.com/gocraft/web"
	pbList "github.com/wizelineacademy/GoWorkshop/proto/list"
	pbUsers "github.com/wizelineacademy/GoWorkshop/proto/users"
)

func (c *Context) home(w web.ResponseWriter, r *web.Request) {
	d := tpl.Data{
		TemplateFile: "home.html",
		Data: struct {
			Error string
		}{},
	}

	d.Render(w, r)
}

func (c *Context) user(w web.ResponseWriter, r *web.Request) {
	var errorMsg string

	id := r.PathParams["id"]
	deleteItemID := r.FormValue("delete_id")
	itemMessage := r.FormValue("item_message")

	if len(deleteItemID) > 0 {
		_, err := c.ListService.DeleteItem(context.Background(), &pbList.DeleteItemRequest{
			Id: deleteItemID,
		})
		if err != nil {
			errorMsg = err.Error()
		}
	}

	if len(itemMessage) > 0 {
		_, err := c.ListService.CreateItem(context.Background(), &pbList.CreateItemRequest{
			Message: itemMessage,
			UserId:  id,
		})
		if err != nil {
			errorMsg = err.Error()
		}
	}

	// gRPC call
	resp, err := c.ListService.GetUserItems(context.Background(), &pbList.GetUserItemsRequest{
		UserId: id,
	})

	if err != nil && len(errorMsg) == 0 {
		errorMsg = err.Error()
	}

	d := tpl.Data{
		TemplateFile: "user.html",
		Data: struct {
			ID    string
			Error string
			Resp  *pbList.GetUserItemsResponse
		}{
			ID:    id,
			Error: errorMsg,
			Resp:  resp,
		},
	}

	d.Render(w, r)
}

func (c *Context) createUser(w web.ResponseWriter, r *web.Request) {
	email := r.FormValue("email")

	// gRPC call
	resp, err := c.UsersService.CreateUser(context.Background(), &pbUsers.CreateUserRequest{
		Email: email,
	})
	if err == nil && resp != nil && resp.Code == http.StatusCreated {
		http.Redirect(w, r.Request, "/user/"+resp.GetId(), http.StatusFound)
		return
	}

	var errorMsg string
	if err != nil {
		errorMsg = err.Error()
	} else if resp != nil {
		errorMsg = resp.Message
	}

	d := tpl.Data{
		TemplateFile: "pages/home.html",
		Data: struct {
			Error string
		}{
			Error: errorMsg,
		},
	}

	d.Render(w, r)
}
