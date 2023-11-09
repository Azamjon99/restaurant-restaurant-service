package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/restaurant/models"
)

type RestaurantRepository interface {
	SaveRestaurantEntity(ctx context.Context, entity *models.RestaurantEntity) error
	UpdateRestaurantEntity(ctx context.Context, entity *models.RestaurantEntity) error
	GetRestaurantEntity(ctx context.Context, entityID string) (*models.RestaurantEntity, error)
	ListRestaurantEntities(ctx context.Context, sort string, page, pageSize int) ([]*models.RestaurantEntity, error)
	DeleteRestaurantEntity(ctx context.Context, entityID string) error
	SaveRestaurant(ctx context.Context, restaurant *models.Restaurant) error
	UpdateRestaurant(ctx context.Context, restaurant *models.Restaurant) error
	GetRestaurant(ctx context.Context, restaurantID string) (*models.Restaurant, error)
	ListRestaurantByEntity(ctx context.Context, entityID string) ([]*models.Restaurant, error)
	DeleteRestaurant(ctx context.Context, restaurantID string) error
	// SearchRestaurantEntities returns restaurant entities sorted by distance asc or desc from the provided location
	SearchRestaurantEntities(ctx context.Context, location []float64, sort string, page, pageSize int) ([]*models.RestaurantEntity, error)
	// SearchRestaurants returns restaurants of a restaurant entity sorted by distance asc from the provided location
	SearchClosestRestaurants(ctx context.Context, entityID string, location []float64) ([]string, error)
}
