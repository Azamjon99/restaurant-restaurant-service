package grpc

import (
	"context"

	pb "github.com/Azamjon99/restaurant-restaurant-service/src/application/protos/restaurant_restaurant"
	"github.com/Azamjon99/restaurant-restaurant-service/src/application/services"
)

type Server struct {
	pb.RestaurantServiceServer
	menuApp       services.MenuAppService
	orderApp      services.OrderAppService
	restaurantApp services.RestaurantAppService
}

func NewServer(menuApp services.MenuAppService, orderApp services.OrderAppService, restaurantApp services.RestaurantAppService) *Server {
	return &Server{
		menuApp:       menuApp,
		orderApp:      orderApp,
		restaurantApp: restaurantApp,
	}
}

func (s *Server) CreateMenu(ctx context.Context, req *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	return s.menuApp.CreateMenu(ctx, req)
}

func (s *Server) UpdateMenu(ctx context.Context, req *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {
	return s.menuApp.UpdateMenu(ctx, req)
}

func (s *Server) DeleteMenu(ctx context.Context, req *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
	return s.menuApp.DeleteMenu(ctx, req)
}

func (s *Server) GetMenu(ctx context.Context, req *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	return s.menuApp.GetMenu(ctx, req)
}

func (s *Server) AddCategory(ctx context.Context, req *pb.AddCategoryRequest) (*pb.AddCategoryResponse, error) {
	return s.menuApp.AddCategory(ctx, req)
}

func (s *Server) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	return s.menuApp.UpdateCategory(ctx, req)
}

func (s *Server) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	return s.menuApp.DeleteCategory(ctx, req)
}

func (s *Server) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error) {
	return s.menuApp.ListCategory(ctx, req)
}

func (s *Server) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	return s.menuApp.AddItem(ctx, req)
}

func (s *Server) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	return s.menuApp.UpdateItem(ctx, req)
}

func (s *Server) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	return s.menuApp.DeleteItem(ctx, req)
}

func (s *Server) ListItem(ctx context.Context, req *pb.ListItemRequest) (*pb.ListItemResponse, error) {
	return s.menuApp.ListItem(ctx, req)
}

func (s *Server) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	return s.orderApp.UpdateOrder(ctx, req)
}

func (s *Server) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	return s.orderApp.DeleteOrder(ctx, req)
}

func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return s.orderApp.GetOrder(ctx, req)
}

func (s *Server) ConfirmOrder(ctx context.Context, req *pb.ConfirmOrderRequest) (*pb.ConfirmOrderResponse, error) {
	return s.orderApp.ConfirmOrder(ctx, req)
}

func (s *Server) DeclineOrder(ctx context.Context, req *pb.DeclineOrderRequest) (*pb.DeclineOrderResponse, error) {
	return s.orderApp.DeclineOrder(ctx, req)
}

func (s *Server) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderResponse, error) {
	return s.orderApp.ListOrder(ctx, req)
}

func (s *Server) CreateRestaurantEntity(ctx context.Context, req *pb.CreateRestaurantEntityRequest) (*pb.CreateRestaurantEntityResponse, error) {
	return s.restaurantApp.CreateRestaurantEntity(ctx, req)
}

func (s *Server) UpdateRestaurantEntity(ctx context.Context, req *pb.UpdateRestaurantEntityRequest) (*pb.UpdateRestaurantEntityResponse, error) {
	return s.restaurantApp.UpdateRestaurantEntity(ctx, req)
}

func (s *Server) ListRestaurantEntity(ctx context.Context, req *pb.ListRestaurantEntityRequest) (*pb.ListRestaurantEntityResponse, error) {
	return s.restaurantApp.ListRestaurantEntity(ctx, req)
}

func (s *Server) DeleteRestaurantEntity(ctx context.Context, req *pb.DeleteRestaurantEntityRequest) (*pb.DeleteRestaurantEntityResponse, error) {
	return s.restaurantApp.DeleteRestaurantEntity(ctx, req)
}

func (s *Server) AddRestaurant(ctx context.Context, req *pb.AddRestaurantRequest) (*pb.AddRestaurantResponse, error) {
	return s.restaurantApp.AddRestaurant(ctx, req)
}

func (s *Server) UpdateRestaurant(ctx context.Context, req *pb.UpdateRestaurantRequest) (*pb.UpdateRestaurantResponse, error) {
	return s.restaurantApp.UpdateRestaurant(ctx, req)
}

func (s *Server) ListRestaurantByEntity(ctx context.Context, req *pb.ListRestaurantByEntityRequest) (*pb.ListRestaurantByEntityResponse, error) {
	return s.restaurantApp.ListRestaurantByEntity(ctx, req)
}

func (s *Server) RemoveRestaurant(ctx context.Context, req *pb.RemoveRestaurantRequest) (*pb.RemoveRestaurantResponse, error) {
	return s.restaurantApp.RemoveRestaurant(ctx, req)
}

func (s *Server) SearchRestaurantEntitiesByLocation(ctx context.Context, req *pb.SearchRestaurantEntitiesByLocationRequest) (*pb.SearchRestaurantEntitiesByLocationResponse, error) {
	return s.restaurantApp.SearchRestaurantEntitiesByLocation(ctx, req)
}

func (s *Server) SearchClosestRestaurants(ctx context.Context, req *pb.SearchClosestRestaurantsRequest) (*pb.SearchClosestRestaurantsResponse, error) {
	return s.restaurantApp.SearchClosestRestaurants(ctx, req)
}
