package handlers

import (
	"log"

	"github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
)

type Service struct{}

func (s *Service) CreateItem(ctx context.Context, in *list.CreateItemRequest) (*list.CreateItemResponse, error) {
	c := shared.DbCollection("list")
	repo := &shared.ListRepository{c}

	itemID, err := repo.Create(&shared.Item{
		Message: in.Message,
		UserId:  in.UserId,
	})

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

func (s *Service) GetUserItems(ctx context.Context, in *list.GetUserItemsRequest) (*list.GetUserItemsResponse, error) {
	items := getUserItems(in.UserId)

	response := &list.GetUserItemsResponse{
		Items: items,
		Code:  200,
	}

	return response, nil
}

func (s *Service) DeleteItem(ctx context.Context, in *list.DeleteItemRequest) (*list.DeleteItemResponse, error) {
	c := shared.DbCollection("list")
	repo := &shared.ListRepository{c}
	err := repo.Delete(in.Id)

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
	c := shared.DbCollection("list")
	repo := &shared.ListRepository{c}
	docs := repo.GetAll(userID)

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
