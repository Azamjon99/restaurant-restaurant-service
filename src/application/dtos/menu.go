package dtos

import (
	"time"

	pb "github.com/Azamjon99/restaurant-restaurant-service/src/application/protos/restaurant_restaurant"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/menu/models"
)

func ToMenuPB(menu *models.Menu) *pb.Menu {

	return &pb.Menu{
		Id:        menu.ID,
		EntityId:  menu.EntityID,
		Name:      menu.Name,
		CreatedAt: menu.CreatedAt.Format(time.RFC3339),
		UpdatedAt: menu.UpdatedAt.Format(time.RFC3339),
	}
}
