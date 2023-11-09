package services

import (
	"context"
	"time"

	"github.com/Azamjon99/restaurant-restaurant-service/src/infrastructure/rand"

	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/restaurant/models"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/restaurant/repositories"
)

type RestaurantService interface {
	CreateRestaurantEntity(ctx context.Context, name string) (*models.RestaurantEntity, error)
	UpdateRestaurantEntity(ctx context.Context, entity *models.RestaurantEntity) error
	GetRestaurantEntity(ctx context.Context, entityID string) (*models.RestaurantEntity, error)
	ListRestaurantEntities(ctx context.Context, sort string, page, pageSize int) ([]*models.RestaurantEntity, error)
	DeleteRestaurantEntity(ctx context.Context, entityID string) error
	AddRestaurant(ctx context.Context, entityID, name string) (*models.Restaurant, error)
	UpdateRestaurant(ctx context.Context, restaurant *models.Restaurant) error
	GetRestaurant(ctx context.Context, restaurantID string) (*models.Restaurant, error)
	ListRestaurantByEntity(ctx context.Context, entity string) ([]*models.Restaurant, error)
	RemoveRestaurant(ctx context.Context, restaurantID string) error
	SearchRestaurantEntitiesByLocation(ctx context.Context, location []float64, sort string, page, pageSize int) ([]*models.RestaurantEntity, error)
	SearchClosestRestaurants(ctx context.Context, entityID string, location []float64) ([]string, error)
}

type restaurantSvcImpl struct {
	restaurantRepo repositories.RestaurantRepository
}

func NewRestaurantSvcImpl(resRepo repositories.RestaurantRepository) RestaurantService {
	return &restaurantSvcImpl{
		restaurantRepo: resRepo,
	}
}

func (s *restaurantSvcImpl) CreateRestaurantEntity(ctx context.Context, name string) (*models.RestaurantEntity, error) {
	entity := models.RestaurantEntity{
		ID:        rand.UUID(),
		Name:      name,
		LogoUrl:   "",
		ImageUrl:  "",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err := s.restaurantRepo.SaveRestaurantEntity(ctx, &entity)

	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (s *restaurantSvcImpl) UpdateRestaurantEntity(ctx context.Context, entity *models.RestaurantEntity) error {

	return s.restaurantRepo.UpdateRestaurantEntity(ctx, entity)
}

func (s *restaurantSvcImpl) GetRestaurantEntity(ctx context.Context, entityID string) (*models.RestaurantEntity, error) {

	return s.restaurantRepo.GetRestaurantEntity(ctx, entityID)
}

func (s *restaurantSvcImpl) ListRestaurantEntities(ctx context.Context, sort string, page, pageSize int) ([]*models.RestaurantEntity, error) {

	return s.restaurantRepo.ListRestaurantEntities(ctx, sort, page, pageSize)
}

func (s *restaurantSvcImpl) DeleteRestaurantEntity(ctx context.Context, entityID string) error {
	return s.restaurantRepo.DeleteRestaurantEntity(ctx, entityID)
}

func (s *restaurantSvcImpl) AddRestaurant(ctx context.Context, entityID, name string) (*models.Restaurant, error) {

	restaurant := models.Restaurant{
		ID:          rand.UUID(),
		EntityID:    entityID,
		Name:        name,
		Description: "",
		Address:     &models.Address{},
		Type:        "",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err := s.restaurantRepo.SaveRestaurant(ctx, &restaurant)

	if err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (s *restaurantSvcImpl) UpdateRestaurant(ctx context.Context, restaurant *models.Restaurant) error {
	return s.restaurantRepo.UpdateRestaurant(ctx, restaurant)
}

func (s *restaurantSvcImpl) GetRestaurant(ctx context.Context, restaurantID string) (*models.Restaurant, error) {
	return s.restaurantRepo.GetRestaurant(ctx, restaurantID)
}

func (s *restaurantSvcImpl) ListRestaurantByEntity(ctx context.Context, entity string) ([]*models.Restaurant, error) {
	return s.restaurantRepo.ListRestaurantByEntity(ctx, entity)
}

func (s *restaurantSvcImpl) RemoveRestaurant(ctx context.Context, restaurantID string) error {
	return s.restaurantRepo.DeleteRestaurant(ctx, restaurantID)
}

func (s *restaurantSvcImpl) SearchRestaurantEntitiesByLocation(ctx context.Context, location []float64, sort string, page, pageSize int) ([]*models.RestaurantEntity, error) {
	return s.restaurantRepo.SearchRestaurantEntities(ctx, location, sort, page, pageSize)
}

func (s *restaurantSvcImpl) SearchClosestRestaurants(ctx context.Context, entityID string, location []float64) ([]string, error) {
	return s.restaurantRepo.SearchClosestRestaurants(ctx, entityID, location)
}
