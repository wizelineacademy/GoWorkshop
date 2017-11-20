package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/wizelineacademy/GoWorkshop/list/models"
	"github.com/wizelineacademy/GoWorkshop/proto/list"
	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) CreateItem(ctx context.Context, in *list.CreateItemRequest) (*list.CreateItemResponse, error) {
	itemID, err := models.CreateItem(in.Message, in.UserId)

	response := new(list.CreateItemResponse)
	if err == nil {
		log.WithField("id", itemID).Info("item created")

		response.Id = itemID
		response.Message = "Item created successfully"
		response.Code = http.StatusCreated
	} else {
		log.WithError(err).Error("unable to create item")

		response.Message = err.Error()
		response.Code = http.StatusInternalServerError
	}

	return response, err
}

func (s *Server) GetUserItems(ctx context.Context, in *list.GetUserItemsRequest) (*list.GetUserItemsResponse, error) {
	items := getUserItems(in.UserId)

	response := &list.GetUserItemsResponse{
		Items: items,
		Code:  http.StatusOK,
	}

	return response, nil
}

func (s *Server) DeleteItem(ctx context.Context, in *list.DeleteItemRequest) (*list.DeleteItemResponse, error) {
	err := models.DeleteItem(in.Id)

	response := new(list.DeleteItemResponse)
	if err == nil {
		log.WithField("id", in.Id).Info("item deleted")

		response.Message = "Item deleted successfully"
		response.Code = http.StatusOK
	} else {
		log.WithError(err).Error("unable to delete item")

		response.Message = err.Error()
		response.Code = http.StatusInternalServerError
	}

	return response, err
}

func getUserItems(userID string) []*list.Item {
	docs := models.GetUserItems(userID)

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
