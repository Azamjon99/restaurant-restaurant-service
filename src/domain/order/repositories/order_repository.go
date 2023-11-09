package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/order/models"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, order *models.Order) error
	UpdateOrder(ctx context.Context, order *models.Order) error
	DeleteOrder(ctx context.Context, orderID string) error
	GetOrder(ctx context.Context, orderID string) (*models.Order, error)
	ConfirmOrder(ctx context.Context, orderID string) error
	DeclineOrder(ctx context.Context, orderID string) error
	ListOrderByRestaurant(ctx context.Context, restaurantID, sort string, page, pageSize int) ([]*models.Order, error)
}
