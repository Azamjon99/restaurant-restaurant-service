package services

import (
	"context"
	"errors"

	"github.com/Azamjon99/restaurant-restaurant-service/src/application/dtos"
	pb "github.com/Azamjon99/restaurant-restaurant-service/src/application/protos/restaurant_restaurant"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/restaurant/services"
)

type RestaurantAppService interface {
	CreateRestaurantEntity(ctx context.Context, req *pb.CreateRestaurantEntityRequest) (*pb.CreateRestaurantEntityResponse, error)
	UpdateRestaurantEntity(ctx context.Context, req *pb.UpdateRestaurantEntityRequest) (*pb.UpdateRestaurantEntityResponse, error)
	ListRestaurantEntity(ctx context.Context, req *pb.ListRestaurantEntityRequest) (*pb.ListRestaurantEntityResponse, error)
	DeleteRestaurantEntity(ctx context.Context, req *pb.DeleteRestaurantEntityRequest) (*pb.DeleteRestaurantEntityResponse, error)
	AddRestaurant(ctx context.Context, req *pb.AddRestaurantRequest) (*pb.AddRestaurantResponse, error)
	UpdateRestaurant(ctx context.Context, req *pb.UpdateRestaurantRequest) (*pb.UpdateRestaurantResponse, error)
	ListRestaurantByEntity(ctx context.Context, req *pb.ListRestaurantByEntityRequest) (*pb.ListRestaurantByEntityResponse, error)
	RemoveRestaurant(ctx context.Context, req *pb.RemoveRestaurantRequest) (*pb.RemoveRestaurantResponse, error)
	SearchRestaurantEntitiesByLocation(ctx context.Context, req *pb.SearchRestaurantEntitiesByLocationRequest) (*pb.SearchRestaurantEntitiesByLocationResponse, error)
	SearchClosestRestaurants(ctx context.Context, req *pb.SearchClosestRestaurantsRequest) (*pb.SearchClosestRestaurantsResponse, error)
}

type restaurantAppImpl struct {
	restaurantSvc services.RestaurantService
}

func NewRestaurantAppService(restaurantSvc services.RestaurantService) RestaurantAppService {
	return &restaurantAppImpl{
		restaurantSvc: restaurantSvc,
	}
}

func (s *restaurantAppImpl) CreateRestaurantEntity(ctx context.Context, req *pb.CreateRestaurantEntityRequest) (*pb.CreateRestaurantEntityResponse, error) {

	if req.GetName() == "" {
		return nil, errors.New("Invalid or empty name")
	}

	entity, err := s.restaurantSvc.CreateRestaurantEntity(ctx, req.GetName())

	if err != nil {
		return nil, err
	}

	return &pb.CreateRestaurantEntityResponse{
		Entity: dtos.ToEntityPB(entity),
	}, nil
}

func (s *restaurantAppImpl) UpdateRestaurantEntity(ctx context.Context, req *pb.UpdateRestaurantEntityRequest) (*pb.UpdateRestaurantEntityResponse, error) {

	if req.GetEntity() == nil {
		return nil, errors.New("Invalid or empty entity")
	}
	if req.GetEntity().GetId() == "" {
		return nil, errors.New("Invalid or missing entity_id")
	}

	entity := dtos.ToEntity(req.GetEntity())

	err := s.restaurantSvc.UpdateRestaurantEntity(ctx, entity)

	if err != nil {
		return nil, err
	}
	return &pb.UpdateRestaurantEntityResponse{
		Entity: dtos.ToEntityPB(entity),
	}, nil
}

func (s *restaurantAppImpl) ListRestaurantEntity(ctx context.Context, req *pb.ListRestaurantEntityRequest) (*pb.ListRestaurantEntityResponse, error) {

	if req.GetPage() == 0 {
		return nil, errors.New("invalid or missing page")
	}
	entities, err := s.restaurantSvc.ListRestaurantEntities(ctx, req.GetSort(), int(req.GetPage()), int(req.GetPageSize()))

	if err != nil {
		return nil, err
	}

	return &pb.ListRestaurantEntityResponse{
		Entities: dtos.ToEntitiesPB(entities),
	}, nil
}

func (s *restaurantAppImpl) DeleteRestaurantEntity(ctx context.Context, req *pb.DeleteRestaurantEntityRequest) (*pb.DeleteRestaurantEntityResponse, error) {

	if req.GetEntityId() == "" {
		return nil, errors.New("Invalid or missing entity")
	}
	err := s.restaurantSvc.DeleteRestaurantEntity(ctx, req.GetEntityId())

	if err != nil {
		return nil, err
	}
	return &pb.DeleteRestaurantEntityResponse{}, nil
}

func (s *restaurantAppImpl) AddRestaurant(ctx context.Context, req *pb.AddRestaurantRequest) (*pb.AddRestaurantResponse, error) {

	if req.GetEntityId() == "" {
		return nil, errors.New("invalid or missing entity_id")
	}
	if req.GetName() == "" {
		return nil, errors.New("invalid or empty name")
	}
	restaurant, err := s.restaurantSvc.AddRestaurant(ctx, req.GetEntityId(), req.GetName())

	if err != nil {
		return nil, err
	}
	return &pb.AddRestaurantResponse{
		Restaurant: dtos.ToRestaurantPB(restaurant),
	}, nil
}

func (s *restaurantAppImpl) UpdateRestaurant(ctx context.Context, req *pb.UpdateRestaurantRequest) (*pb.UpdateRestaurantResponse, error) {

	if req.GetRestaurant() == nil {
		return nil, errors.New("Invalid or empty restaurant")
	}
	restaurant := dtos.ToRestaurant(req.GetRestaurant())

	err := s.restaurantSvc.UpdateRestaurant(ctx, restaurant)

	if err != nil {
		return nil, err
	}
	return &pb.UpdateRestaurantResponse{
		Restaurant: dtos.ToRestaurantPB(restaurant),
	}, nil
}

func (s *restaurantAppImpl) ListRestaurantByEntity(ctx context.Context, req *pb.ListRestaurantByEntityRequest) (*pb.ListRestaurantByEntityResponse, error) {

	if req.GetEntityId() == "" {
		return nil, errors.New("Invalid or missing entity_id")
	}
	restaurants, err := s.restaurantSvc.ListRestaurantByEntity(ctx, req.GetEntityId())

	if err != nil {
		return nil, err
	}
	return &pb.ListRestaurantByEntityResponse{
		Restaurants: dtos.ToRestaurantsDB(restaurants),
	}, nil
}

func (s *restaurantAppImpl) RemoveRestaurant(ctx context.Context, req *pb.RemoveRestaurantRequest) (*pb.RemoveRestaurantResponse, error) {

	if req.GetRestaurantId() == "" {
		return nil, errors.New("Invalid or missing restaurant-id")
	}
	err := s.restaurantSvc.RemoveRestaurant(ctx, req.GetRestaurantId())

	if err != nil {
		return nil, err
	}
	return &pb.RemoveRestaurantResponse{}, nil
}

func (s *restaurantAppImpl) SearchRestaurantEntitiesByLocation(ctx context.Context, req *pb.SearchRestaurantEntitiesByLocationRequest) (*pb.SearchRestaurantEntitiesByLocationResponse, error) {

	entites, err := s.restaurantSvc.SearchRestaurantEntitiesByLocation(ctx, req.GetLocation(), req.GetSort(), int(req.GetPage()), int(req.GetPageSize()))

	if err != nil {
		return nil, err
	}
	return &pb.SearchRestaurantEntitiesByLocationResponse{
		Entities: dtos.ToEntitiesPB(entites),
	}, nil
}

func (s *restaurantAppImpl) SearchClosestRestaurants(ctx context.Context, req *pb.SearchClosestRestaurantsRequest) (*pb.SearchClosestRestaurantsResponse, error) {

	if req.GetEntityId() == "" {
		return nil, errors.New("Invalid or missing entity_id")
	}
	ids, err := s.restaurantSvc.SearchClosestRestaurants(ctx, req.GetEntityId(), req.GetLocation())

	if err != nil {
		return nil, err
	}
	return &pb.SearchClosestRestaurantsResponse{
		RestaurantIds: ids,
	}, nil
}
