package server

import (
	"log"

	"github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) CreateItem(ctx context.Context, in *list.CreateItemRequest) (*list.CreateItemResponse, error) {
	itemID, err := shared.CreateItem(in.Message, in.UserId)

	response := new(list.CreateItemResponse)
	if err == nil {
		log.Printf("[item.Create] New item ID: %s", itemID)

		response.Id = itemID
		response.Message = "Item created successfully"
		response.Code = 200
	} else {
		response.Message = err.Error()
		response.Code = 500
	}

	return response, err
}

func (s *Server) GetUserItems(ctx context.Context, in *list.GetUserItemsRequest) (*list.GetUserItemsResponse, error) {
	items := getUserItems(in.UserId)

	response := &list.GetUserItemsResponse{
		Items: items,
		Code:  200,
	}

	return response, nil
}

func (s *Server) DeleteItem(ctx context.Context, in *list.DeleteItemRequest) (*list.DeleteItemResponse, error) {
	err := shared.DeleteItem(in.Id)

	response := new(list.DeleteItemResponse)
	if err == nil {
		log.Printf("[item.Delete] Deleted item ID: %s", in.Id)

		response.Message = "Item deleted successfully"
		response.Code = 200
	} else {
		response.Message = err.Error()
		response.Code = 500
	}

	return response, err
}

func getUserItems(userID string) []*list.Item {
	docs := shared.GetUserItems(userID)

	items := []*list.Item{}
	for _, item := range docs {
		items = append(items, &list.Item{
			Id:      item.Id.Hex(),
			Message: item.Message,
			UserId:  item.UserId,
		})
	}

	return items
}
