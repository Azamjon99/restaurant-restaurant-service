package services

import (
	"context"

	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/order/models"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/order/repositories"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	UpdateOrder(ctx context.Context, order *models.Order) error
	DeleteOrder(ctx context.Context, orderID string) error
	GetOrder(ctx context.Context, orderID string) (*models.Order, error)
	ConfirmOrder(ctx context.Context, orderID string) error
	DeclineOrder(ctx context.Context, orderID string) error
	ListOrderByRestaurant(ctx context.Context, restaurantID, sort string, page, pageSize int) ([]*models.Order, error)
}

type orderSvcImpl struct {
	orderRepo repositories.OrderRepository
}

func NewOrderSvcImpl(orderRepo repositories.OrderRepository) OrderService {
	return &orderSvcImpl{
		orderRepo: orderRepo,
	}
}

func (s *orderSvcImpl) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {

	err := s.orderRepo.SaveOrder(ctx, order)

	return order, err
}
func (s *orderSvcImpl) UpdateOrder(ctx context.Context, order *models.Order) error {

	return s.orderRepo.UpdateOrder(ctx, order)
}
func (s *orderSvcImpl) DeleteOrder(ctx context.Context, orderID string) error {

	return s.orderRepo.DeleteOrder(ctx, orderID)
}
func (s *orderSvcImpl) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {

	return s.orderRepo.GetOrder(ctx, orderID)
}
func (s *orderSvcImpl) ConfirmOrder(ctx context.Context, orderID string) error {

	return s.orderRepo.ConfirmOrder(ctx, orderID)
}
func (s *orderSvcImpl) DeclineOrder(ctx context.Context, orderID string) error {

	return s.orderRepo.DeclineOrder(ctx, orderID)
}
func (s *orderSvcImpl) ListOrderByRestaurant(ctx context.Context, restaurantID, sort string, page, pageSize int) ([]*models.Order, error) {

	return s.orderRepo.ListOrderByRestaurant(ctx, restaurantID, sort, page, pageSize)
}
