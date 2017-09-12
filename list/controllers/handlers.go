package controllers

import (
	"github.com/wizelineacademy/GoWorkshop/list/models"
	pb "github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/shared/config"
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

	appContext := config.NewContext()
	defer appContext.Close()

	c := appContext.DbCollection("list")
	repo := &models.ListRepository{
		C: c,
	}
	err = repo.Create(item)

	response = new(pb.CreateItemResponse)
	if err == nil {
		response.Message = "Item created successfully"
		response.Code = 200
	} else {
		response.Message = err.Error()
		response.Code = 500
	}

	return response, err
}

// GetAllItems implementation
func (s *Service) GetAllItems(ctx context.Context, in *pb.GetAllItemsRequest) (response *pb.GetAllItemsResponse, err error) {
	appContext := config.NewContext()
	defer appContext.Close()

	c := appContext.DbCollection("list")
	repo := &models.ListRepository{
		C: c,
	}
	items := repo.GetAll()

	pbItems := []*pb.Item{}
	for _, item := range items {
		pbItems = append(pbItems, &pb.Item{
			Id:      item.Id.Hex(),
			Message: item.Message,
			UserId:  item.UserId,
		})
	}
	response = &pb.GetAllItemsResponse{
		Items: pbItems,
		Code:  200,
	}

	return response, err
}

// DeleteItem implementation
func (s *Service) DeleteItem(ctx context.Context, in *pb.DeleteItemRequest) (response *pb.DeleteItemResponse, err error) {
	appContext := config.NewContext()
	defer appContext.Close()

	c := appContext.DbCollection("list")
	repo := &models.ListRepository{
		C: c,
	}
	err = repo.Delete(in.Id)

	if err == nil {
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

	return response, err
}
