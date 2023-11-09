package dtos

import (
	"time"

	pb "github.com/Azamjon99/restaurant-restaurant-service/src/application/protos/restaurant_restaurant"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/menu/models"
)

func ToCategoryPB(category *models.Category) *pb.Category {

	return &pb.Category{
		Id:          category.ID,
		MenuId:      category.MenuID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   category.UpdatedAt.Format(time.RFC3339),
	}
}
func ToCategoriesPB(categories []*models.Category) []*pb.Category {

	result := make([]*pb.Category, len(categories))
	for i := range categories {
		result[i] = ToCategoryPB(categories[i])
	}
	return result
}
