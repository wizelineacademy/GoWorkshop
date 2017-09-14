package controllers

import (
	"log"

	"github.com/wizelineacademy/GoWorkshop/list/models"
	pb "github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"golang.org/x/net/context"
)

// Service struct
type Service struct{}

// CreateItem implementation
func (s *Service) CreateItem(ctx context.Context, in *pb.CreateItemRequest) (response *pb.CreateItemResponse, err error) {
	item := &models.Item{
		Message: in.Message,
		UserId:  in.UserId,
	}

	appContext := shared.NewContext()
	defer appContext.Close()

	c := appContext.DbCollection("list")
	repo := &models.ListRepository{
		C: c,
	}

	var itemID string
	itemID, err = repo.Create(item)

	response = new(pb.CreateItemResponse)
	if err == nil {
		log.Printf("[item.Create] New item ID: %s", itemID)

		response.Id = itemID
		response.Message = "Item created successfully"
		response.Code = 200
	} else {
		response.Message = err.Error()
		response.Code = 500
	}

	return
}

// GetUserItems implementation
func (s *Service) GetUserItems(ctx context.Context, in *pb.GetUserItemsRequest) (response *pb.GetUserItemsResponse, err error) {
	appContext := shared.NewContext()
	defer appContext.Close()

	c := appContext.DbCollection("list")
	repo := &models.ListRepository{
		C: c,
	}
	items := repo.GetAll(in.UserId)

	pbItems := []*pb.Item{}
	for _, item := range items {
		pbItems = append(pbItems, &pb.Item{
			Id:      item.Id.Hex(),
			Message: item.Message,
			UserId:  item.UserId,
		})
	}
	response = &pb.GetUserItemsResponse{
		Items: pbItems,
		Code:  200,
	}

	return
}

// DeleteItem implementation
func (s *Service) DeleteItem(ctx context.Context, in *pb.DeleteItemRequest) (response *pb.DeleteItemResponse, err error) {
	appContext := shared.NewContext()
	defer appContext.Close()

	c := appContext.DbCollection("list")
	repo := &models.ListRepository{
		C: c,
	}
	err = repo.Delete(in.Id)

	if err == nil {
		log.Printf("[item.Delete] Deleted item ID: %s", in.Id)

		response = &pb.DeleteItemResponse{
			Message: "Item deleted successfully",
			Code:    200,
		}
	} else {
		response = &pb.DeleteItemResponse{
			Message: err.Error(),
			Code:    500,
		}
	}

	return
}
