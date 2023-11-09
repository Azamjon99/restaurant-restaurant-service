package dtos

import (
	"time"

	pb "github.com/Azamjon99/restaurant-restaurant-service/src/application/protos/restaurant_restaurant"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/menu/models"
)

func ToItemPB(item *models.Item) *pb.Item {

	return &pb.Item{
		Id:          item.ID,
		CategoryId:  item.CategoryID,
		Name:        item.Name,
		Description: item.Description,
		ImageUrl:    item.ImageUrl,
		Price:       int32(item.Price),
		CreatedAt:   item.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   item.UpdatedAt.Format(time.RFC3339),
	}
}
func ToItemsPB(items []*models.Item) []*pb.Item {

	result := make([]*pb.Item, len(items))
	for i := range items {
		result[i] = ToItemPB(items[i])
	}
	return result
}
func ToItem(item *pb.Item) *models.Item {

	return &models.Item{
		ID:          item.GetId(),
		CategoryID:  item.GetCategoryId(),
		Name:        item.GetName(),
		Description: item.GetDescription(),
		ImageUrl:    item.GetImageUrl(),
		Price:       int(item.GetPrice()),
		CreatedAt:   toTime(item.GetCreatedAt()),
		UpdatedAt:   toTime(item.GetUpdatedAt()),
	}
}
