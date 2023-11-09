package services

import (
	"context"
	"errors"

	"github.com/Azamjon99/restaurant-restaurant-service/src/application/dtos"
	pb "github.com/Azamjon99/restaurant-restaurant-service/src/application/protos/restaurant_restaurant"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/order/services"
)

type OrderAppService interface {
	UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error)
	DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error)
	GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error)
	ConfirmOrder(ctx context.Context, req *pb.ConfirmOrderRequest) (*pb.ConfirmOrderResponse, error)
	DeclineOrder(ctx context.Context, req *pb.DeclineOrderRequest) (*pb.DeclineOrderResponse, error)
	ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderResponse, error)
}

type orderAppImpl struct {
	orderSvc services.OrderService
}

func NewOrderAppService(orderSvc services.OrderService) OrderAppService {
	return &orderAppImpl{
		orderSvc: orderSvc,
	}
}

func (s *orderAppImpl) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {

	if req.GetOrder() == nil {
		return nil, errors.New("Invalid or empty order")
	}
	if req.GetOrder().GetRestaurantId() == "" {
		return nil, errors.New("Invalid or missing restaurant_id")
	}

	order := dtos.ToOrder(req.GetOrder())

	err := s.orderSvc.UpdateOrder(ctx, order)

	if err != nil {
		return nil, err
	}
	return &pb.UpdateOrderResponse{
		Order: dtos.ToOrderPB(order),
	}, nil
}

func (s *orderAppImpl) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {

	if req.GetOrderId() == "" {
		return nil, errors.New("Invalid or missing order_id")
	}

	err := s.orderSvc.DeleteOrder(ctx, req.GetOrderId())

	if err != nil {
		return nil, err
	}

	return &pb.DeleteOrderResponse{}, nil
}

func (s *orderAppImpl) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {

	if req.GetOrderId() == "" {

		return nil, errors.New("Invalid or missing order_id")
	}
	order, err := s.orderSvc.GetOrder(ctx, req.GetOrderId())

	if err != nil {
		return nil, err
	}
	return &pb.GetOrderResponse{
		Order: dtos.ToOrderPB(order),
	}, nil
}

func (s *orderAppImpl) ConfirmOrder(ctx context.Context, req *pb.ConfirmOrderRequest) (*pb.ConfirmOrderResponse, error) {

	if req.GetOrderId() == "" {
		return nil, errors.New("Invalid or missing order_id")
	}

	err := s.orderSvc.ConfirmOrder(ctx, req.GetOrderId())

	if err != nil {
		return nil, err
	}

	return &pb.ConfirmOrderResponse{}, nil
}

func (s *orderAppImpl) DeclineOrder(ctx context.Context, req *pb.DeclineOrderRequest) (*pb.DeclineOrderResponse, error) {

	if req.GetOrderId() == "" {
		return nil, errors.New("Invalid or missing order_id")
	}

	err := s.orderSvc.DeclineOrder(ctx, req.GetOrderId())

	if err != nil {
		return nil, err
	}

	return &pb.DeclineOrderResponse{}, nil
}

func (s *orderAppImpl) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderResponse, error) {

	if req.GetRestaurantId() == "" {
		return nil, errors.New("Invalid or missing restaurant_id")
	}

	orders, err := s.orderSvc.ListOrderByRestaurant(ctx, req.RestaurantId, req.GetSort(), int(req.GetPage()), int(req.GetPageSize()))

	if err != nil {
		return nil, err
	}

	return &pb.ListOrderResponse{
		Orders: dtos.ToOrdersPB(orders),
	}, nil
}
