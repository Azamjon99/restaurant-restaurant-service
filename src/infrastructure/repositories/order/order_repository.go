package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/order/models"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/order/repositories"
	"github.com/Azamjon99/restaurant-restaurant-service/src/infrastructure/repositories/utils"
	"gorm.io/gorm"
)

type orderRepoImpl struct {
	db *gorm.DB
}

func NewOrderRepoImpl(db *gorm.DB) repositories.OrderRepository {
	return &orderRepoImpl{
		db: db,
	}
}

const tableOrder = "restaurant.orders"

func (r *orderRepoImpl) SaveOrder(ctx context.Context, order *models.Order) error {

	result := r.db.WithContext(ctx).Table(tableOrder).Create(order)

	return result.Error
}

func (r *orderRepoImpl) UpdateOrder(ctx context.Context, order *models.Order) error {

	result := r.db.WithContext(ctx).Table(tableOrder).Save(order)

	return result.Error
}

func (r *orderRepoImpl) DeleteOrder(ctx context.Context, orderID string) error {

	var order models.Order

	result := r.db.WithContext(ctx).Table(tableOrder).Delete(&order, "id = ?", orderID)

	return result.Error
}

func (r *orderRepoImpl) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {

	var order models.Order

	result := r.db.WithContext(ctx).Table(tableOrder).First(&order, "id = ?", orderID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (r *orderRepoImpl) ConfirmOrder(ctx context.Context, orderID string) error {

	result := r.db.WithContext(ctx).Table(tableOrder).Where("id = ?", orderID).Update("status", models.OrderStatusConfirmed)

	return result.Error
}

func (r *orderRepoImpl) DeclineOrder(ctx context.Context, orderID string) error {

	result := r.db.WithContext(ctx).Table(tableOrder).Where("id = ?", orderID).Update("status", models.OrderStatusDeclined)

	return result.Error
}

func (r *orderRepoImpl) ListOrderByRestaurant(ctx context.Context, restaurantID, sort string, page, pageSize int) ([]*models.Order, error) {

	var orders []*models.Order

	result := r.db.WithContext(ctx).Table(tableOrder).Where("restaurant_id = ?", restaurantID).
		Scopes(utils.Sort(sort), utils.Paginate(page, pageSize)).
		Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}