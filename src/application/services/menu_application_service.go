package services

import (
	"context"
	"errors"

	"github.com/Azamjon99/restaurant-restaurant-service/src/application/dtos"
	pb "github.com/Azamjon99/restaurant-restaurant-service/src/application/protos/restaurant_restaurant"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/menu/models"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/menu/services"
)

type MenuAppService interface {
	CreateMenu(ctx context.Context, req *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error)
	UpdateMenu(ctx context.Context, req *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error)
	DeleteMenu(ctx context.Context, req *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error)
	GetMenu(ctx context.Context, req *pb.GetMenuRequest) (*pb.GetMenuResponse, error)
	AddCategory(ctx context.Context, req *pb.AddCategoryRequest) (*pb.AddCategoryResponse, error)
	UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error)
	DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error)
	ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error)
	AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error)
	UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error)
	DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error)
	ListItem(ctx context.Context, req *pb.ListItemRequest) (*pb.ListItemResponse, error)
}

type menuAppSvcImpl struct {
	menuSvc services.MenuService
}

func NewMenuAppService(menuSvc services.MenuService) MenuAppService {
	return &menuAppSvcImpl{
		menuSvc: menuSvc,
	}
}

func (s *menuAppSvcImpl) CreateMenu(ctx context.Context, req *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {

	if req.GetEntityId() == "" {
		return nil, errors.New("Invalid or missing entity_id")
	}

	if req.GetName() == "" {

		return nil, errors.New("Invalid or missing name")
	}

	menu, err := s.menuSvc.CreateMenu(ctx, req.GetEntityId(), req.GetName())

	if err != nil {
		return nil, err
	}

	return &pb.CreateMenuResponse{
		Menu: dtos.ToMenuPB(menu),
	}, nil
}

func (s *menuAppSvcImpl) UpdateMenu(ctx context.Context, req *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {

	if req.GetMenuId() == "" {
		return nil, errors.New("Invalid or missing entity_id")
	}

	if req.GetName() == "" {

		return nil, errors.New("Invalid or missing name")
	}

	err := s.menuSvc.UpdateMenu(ctx, req.GetMenuId(), req.GetName())
	if err != nil {
		return nil, err
	}
	return &pb.UpdateMenuResponse{}, nil
}

func (s *menuAppSvcImpl) DeleteMenu(ctx context.Context, req *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {

	if req.GetMenuId() == "" {
		return nil, errors.New("Invalid or missing entity_id")
	}

	err := s.menuSvc.DeleteMenu(ctx, req.GetMenuId())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteMenuResponse{}, nil
}

func (s *menuAppSvcImpl) GetMenu(ctx context.Context, req *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {

	var (
		menu *models.Menu
		err  error
	)

	if req.GetEntityId() != "" {

		menu, err = s.menuSvc.GetMenuByEntity(ctx, req.GetEntityId())
	} else if req.GetMenuId() != "" {
		menu, err = s.menuSvc.GetMenu(ctx, req.GetMenuId())
	}
	if err != nil {
		return nil, err
	}
	return &pb.GetMenuResponse{
		Menu: dtos.ToMenuPB(menu),
	}, nil
}

func (s *menuAppSvcImpl) AddCategory(ctx context.Context, req *pb.AddCategoryRequest) (*pb.AddCategoryResponse, error) {

	if req.GetMenuId() == "" {
		return nil, errors.New("Invalid or missing menu_id")
	}
	if req.GetName() == "" {
		return nil, errors.New("Invalid or empty name")
	}

	category, err := s.menuSvc.AddCategory(ctx, req.GetMenuId(), req.GetName())

	if err != nil {
		return nil, err
	}

	return &pb.AddCategoryResponse{
		Category: dtos.ToCategoryPB(category),
	}, nil
}

func (s *menuAppSvcImpl) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {

	if req.GetCategoryId() == "" {
		return nil, errors.New("Invalid or missing category_id")
	}
	if req.GetName() == "" {
		return nil, errors.New("Invalid or empty name")
	}

	err := s.menuSvc.UpdateCategory(ctx, req.GetCategoryId(), req.GetName())

	if err != nil {
		return nil, err
	}
	return &pb.UpdateCategoryResponse{}, nil
}

func (s *menuAppSvcImpl) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {

	if req.GetCategoryId() == "" {
		return nil, errors.New("Invalid or missing category_id")
	}

	err := s.menuSvc.DeleteCategory(ctx, req.GetCategoryId())

	if err != nil {
		return nil, err
	}

	return &pb.DeleteCategoryResponse{}, nil
}

func (s *menuAppSvcImpl) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error) {

	if req.GetMenuId() == "" {
		return nil, errors.New("Invalid or missing menu_id")
	}

	categories, err := s.menuSvc.ListCategoryByMenu(ctx, req.GetMenuId())

	if err != nil {
		return nil, err
	}

	return &pb.ListCategoryResponse{
		Categories: dtos.ToCategoriesPB(categories),
	}, nil
}

func (s *menuAppSvcImpl) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {

	if req.GetCategoryId() == "" {
		return nil, errors.New("Invalid or missing category_id")
	}
	if req.GetName() == "" {
		return nil, errors.New("Invalid or empty name")
	}
	item, err := s.menuSvc.AddItem(ctx, req.GetCategoryId(), req.GetName())

	if err != nil {
		return nil, err
	}

	return &pb.AddItemResponse{
		Item: dtos.ToItemPB(item),
	}, nil
}

func (s *menuAppSvcImpl) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {

	if req.GetItem() == nil {
		return nil, errors.New("Invalid or empty item")
	}
	if req.GetItem().GetId() == "" {
		return nil, errors.New("Invalid or missing item_id")
	}
	if req.GetItem().GetName() == "" {
		return nil, errors.New("Invalid or missing name")
	}
	item := dtos.ToItem(req.GetItem())

	err := s.menuSvc.UpdateItem(ctx, item)

	if err != nil {
		return nil, err
	}
	return &pb.UpdateItemResponse{}, nil
}

func (s *menuAppSvcImpl) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {

	if req.GetItemId() == "" {
		return nil, errors.New("Invalid or missing item_id")
	}
	err := s.menuSvc.DeleteItem(ctx, req.GetItemId())

	if err != nil {
		return nil, err
	}

	return &pb.DeleteItemResponse{}, nil
}

func (s *menuAppSvcImpl) ListItem(ctx context.Context, req *pb.ListItemRequest) (*pb.ListItemResponse, error) {

	var (
		items []*models.Item
		err   error
	)
	if req.GetCategoryId() != "" {
		items, err = s.menuSvc.ListItemByCategory(ctx, req.GetCategoryId())
	} else if req.GetItemIds() != nil {
		items, err = s.menuSvc.ListItemByIDs(ctx, req.GetItemIds())
	}

	if err != nil {
		return nil, err
	}

	return &pb.ListItemResponse{
		Items: dtos.ToItemsPB(items),
	}, nil
}
