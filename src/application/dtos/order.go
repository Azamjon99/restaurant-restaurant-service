package dtos

import (
	"time"

	pb "github.com/Azamjon99/restaurant-restaurant-service/src/application/protos/restaurant_restaurant"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/order/models"
	"github.com/Azamjon99/restaurant-restaurant-service/src/infrastructure/rand"
)

func ToOrderPB(order *models.Order) *pb.Order {

	return &pb.Order{
		Id:           order.ID,
		RestaurantId: order.RestaurantID,
		Status:       order.Status,
		Eater: &pb.Order_EaterInfo{
			Id:   order.Eater.ID,
			Name: order.Eater.Name,
		},
		Delivery: &pb.Order_DeliveryInfo{
			Name:      order.Delivery.Name,
			Longitude: order.Delivery.Longitude,
			Latitude:  order.Delivery.Latitude,
		},
		Payment: &pb.Order_PaymentInfo{
			Method:        order.Payment.Method,
			CardId:        order.Payment.CardID,
			DeliveryPrice: int32(order.Payment.DeliveryFee),
			TotalPrice:    int32(order.Payment.TotalPrice),
		},
		Items:     toOrderItemsPB(order.Items),
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
	}
}
func ToOrdersPB(orders []*models.Order) []*pb.Order {
	result := make([]*pb.Order, len(orders))
	for i := range orders {
		result[i] = ToOrderPB(orders[i])
	}
	return result
}

func toOrderItemPB(item *models.OrderItem) *pb.OrderItem {

	return &pb.OrderItem{
		Id:         item.ID,
		Name:       item.Name,
		Quantity:   int32(item.Quantity),
		Price:      int32(item.Price),
		TotalPrice: int32(item.TotalPrice),
	}
}
func toOrderItemsPB(items []*models.OrderItem) []*pb.OrderItem {

	result := make([]*pb.OrderItem, len(items))
	for i := range items {
		result[i] = toOrderItemPB(items[i])
	}
	return result
}
func ToOrder(order *pb.Order) *models.Order {

	return &models.Order{
		ID:           rand.UUID(),
		RestaurantID: order.GetRestaurantId(),
		Status:       order.GetStatus(),
		Eater: &models.EaterInfo{
			ID:   order.GetEater().GetId(),
			Name: order.GetEater().GetName(),
		},
		Delivery: &models.DeliveryInfo{
			Name:      order.GetDelivery().GetName(),
			Longitude: order.GetDelivery().GetLongitude(),
			Latitude:  order.GetDelivery().GetLatitude(),
		},
		Payment: &models.PaymentInfo{
			Method:      order.GetPayment().GetMethod(),
			DeliveryFee: int(order.GetPayment().GetDeliveryPrice()),
			TotalPrice:  int(order.GetPayment().GetTotalPrice()),
			CardID:      order.GetPayment().GetCardId(),
		},
		Items:     toOrderItems(order.GetItems()),
		CreatedAt: toTime(order.GetCreatedAt()),
		UpdatedAt: toTime(order.GetUpdatedAt()),
	}
}

func toOrderItem(item *pb.OrderItem) *models.OrderItem {

	return &models.OrderItem{
		ID:         item.GetId(),
		Name:       item.GetName(),
		Price:      int(item.GetPrice()),
		Quantity:   int(item.GetQuantity()),
		TotalPrice: int(item.GetPrice()) * int(item.GetQuantity()),
	}

}
func toOrderItems(items []*pb.OrderItem) []*models.OrderItem {

	result := make([]*models.OrderItem, len(items))
	for i := range items {
		result[i] = toOrderItem(items[i])
	}
	return result
}
