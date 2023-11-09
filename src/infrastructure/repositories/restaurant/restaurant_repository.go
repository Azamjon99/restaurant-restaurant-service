package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/restaurant/models"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/restaurant/repositories"
	"github.com/Azamjon99/restaurant-restaurant-service/src/infrastructure/repositories/utils"
	"gorm.io/gorm"
)

type restaurantRepoImpl struct {
	db *gorm.DB
}

func NewRestaurantRepoImpl(db *gorm.DB) repositories.RestaurantRepository {
	return &restaurantRepoImpl{
		db: db,
	}
}

const (
	tableRestaurantEntity = "restaurant.restaurant_entities"
	tableRestaurant       = "restaurant.restaurants"
)

func (r *restaurantRepoImpl) SaveRestaurantEntity(ctx context.Context, entity *models.RestaurantEntity) error {

	result := r.db.WithContext(ctx).Table(tableRestaurantEntity).Create(entity)

	return result.Error
}

func (r *restaurantRepoImpl) UpdateRestaurantEntity(ctx context.Context, entity *models.RestaurantEntity) error {

	result := r.db.WithContext(ctx).Table(tableRestaurantEntity).Save(entity)

	return result.Error
}

func (r *restaurantRepoImpl) GetRestaurantEntity(ctx context.Context, entityID string) (*models.RestaurantEntity, error) {

	var entity models.RestaurantEntity
	result := r.db.WithContext(ctx).Table(tableRestaurantEntity).First(&entity, "id = ?", entityID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

func (r *restaurantRepoImpl) ListRestaurantEntities(ctx context.Context, sort string, page, pageSize int) ([]*models.RestaurantEntity, error) {

	var entities []*models.RestaurantEntity

	result := r.db.WithContext(ctx).Table(tableRestaurantEntity).
		Scopes(utils.Sort(sort), utils.Paginate(page, pageSize)).
		Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func (r *restaurantRepoImpl) DeleteRestaurantEntity(ctx context.Context, entityID string) error {

	var entity models.RestaurantEntity
	result := r.db.WithContext(ctx).Table(tableRestaurantEntity).Delete(&entity, "id = ?", entityID)
	return result.Error
}

func (r *restaurantRepoImpl) SaveRestaurant(ctx context.Context, restaurant *models.Restaurant) error {

	result := r.db.WithContext(ctx).Table(tableRestaurant).Create(restaurant)
	return result.Error
}

func (r *restaurantRepoImpl) UpdateRestaurant(ctx context.Context, restaurant *models.Restaurant) error {

	result := r.db.WithContext(ctx).Table(tableRestaurant).Save(restaurant)
	return result.Error
}

func (r *restaurantRepoImpl) GetRestaurant(ctx context.Context, restaurantID string) (*models.Restaurant, error) {

	var restaurant models.Restaurant
	result := r.db.WithContext(ctx).Table(tableRestaurant).First(&restaurant, "id = ?", restaurantID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &restaurant, nil
}

func (r *restaurantRepoImpl) ListRestaurantByEntity(ctx context.Context, entityID string) ([]*models.Restaurant, error) {

	var restaurants []*models.Restaurant

	result := r.db.WithContext(ctx).Table(tableRestaurant).Find(&restaurants, "entity_id = ?", entityID)

	if result.Error != nil {
		return nil, result.Error
	}
	return restaurants, nil
}

func (r *restaurantRepoImpl) DeleteRestaurant(ctx context.Context, restaurantID string) error {

	var restaurant models.Restaurant
	result := r.db.WithContext(ctx).Table(tableRestaurant).Delete(&restaurant, "id = ?", restaurantID)

	return result.Error
}

func (r *restaurantRepoImpl) SearchRestaurantEntities(ctx context.Context, location []float64, sort string, page, pageSize int) ([]*models.RestaurantEntity, error) {

	return nil, nil
}

func (r *restaurantRepoImpl) SearchClosestRestaurants(ctx context.Context, entityID string, location []float64) ([]string, error) {

	return nil, nil
}
