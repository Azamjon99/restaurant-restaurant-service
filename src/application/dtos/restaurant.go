package dtos

import (
	"time"

	pb "github.com/Azamjon99/restaurant-restaurant-service/src/application/protos/restaurant_restaurant"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/restaurant/models"
)

func ToEntityPB(entity *models.RestaurantEntity) *pb.RestaurantEntity {

	return &pb.RestaurantEntity{
		Id:        entity.ID,
		Name:      entity.Name,
		LogoUrl:   entity.LogoUrl,
		ImageUrl:  entity.ImageUrl,
		CreatedAt: entity.CreatedAt.Format(time.RFC3339),
		UpdatedAt: entity.UpdatedAt.Format(time.RFC3339),
	}
}
func ToEntitiesPB(entities []*models.RestaurantEntity) []*pb.RestaurantEntity {

	result := make([]*pb.RestaurantEntity, len(entities))

	for i := range entities {
		result[i] = ToEntityPB(entities[i])
	}
	return result
}
func ToEntity(entity *pb.RestaurantEntity) *models.RestaurantEntity {
	return &models.RestaurantEntity{
		ID:        entity.GetId(),
		Name:      entity.GetName(),
		LogoUrl:   entity.GetLogoUrl(),
		ImageUrl:  entity.GetImageUrl(),
		CreatedAt: toTime(entity.GetCreatedAt()),
		UpdatedAt: toTime(entity.GetUpdatedAt()),
	}
}
func toAddressPB(address *models.Address) *pb.Address {

	if address == nil {
		return nil
	}
	return &pb.Address{
		Name:        address.Name,
		Description: address.Description,
		Longitude:   address.Longitude,
		Latitude:    address.Latitude,
	}
}
func toAddress(address *pb.Address) *models.Address {

	if address == nil {
		return nil
	}
	return &models.Address{
		Name:        address.GetName(),
		Description: address.GetDescription(),
		Longitude:   address.GetLongitude(),
		Latitude:    address.GetLatitude(),
	}
}
func ToRestaurantPB(restaurant *models.Restaurant) *pb.Restaurant {
	return &pb.Restaurant{
		Id:          restaurant.ID,
		EntityId:    restaurant.EntityID,
		Name:        restaurant.Name,
		Description: restaurant.Description,
		Address:     toAddressPB(restaurant.Address),
		Type:        restaurant.Type,
		CreatedAt:   restaurant.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   restaurant.UpdatedAt.Format(time.RFC3339),
	}
}

func ToRestaurantsDB(restaurants []*models.Restaurant) []*pb.Restaurant {

	result := make([]*pb.Restaurant, len(restaurants))

	for i := range restaurants {
		result[i] = ToRestaurantPB(restaurants[i])
	}
	return result
}
func ToRestaurant(restaurant *pb.Restaurant) *models.Restaurant {

	return &models.Restaurant{
		ID:          restaurant.GetId(),
		EntityID:    restaurant.GetEntityId(),
		Name:        restaurant.GetName(),
		Description: restaurant.GetDescription(),
		Address:     toAddress(restaurant.GetAddress()),
		Type:        restaurant.GetType(),
		CreatedAt:   toTime(restaurant.GetCreatedAt()),
		UpdatedAt:   toTime(restaurant.GetUpdatedAt()),
	}
}
