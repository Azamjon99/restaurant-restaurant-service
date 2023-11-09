package services

import (
	"context"
	"time"

	"github.com/Azamjon99/restaurant-restaurant-service/src/infrastructure/rand"

	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/menu/models"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/menu/repositories"
)

type MenuService interface {
	CreateMenu(ctx context.Context, entityID, name string) (*models.Menu, error)
	UpdateMenu(ctx context.Context, menu_ID, name string) error
	DeleteMenu(ctx context.Context, menuID string) error
	GetMenu(ctx context.Context, menuID string) (*models.Menu, error)
	GetMenuByEntity(ctx context.Context, entityID string) (*models.Menu, error)
	AddCategory(ctx context.Context, menuID, name string) (*models.Category, error)
	UpdateCategory(ctx context.Context, categoryID, name string) error
	DeleteCategory(ctx context.Context, categoryID string) error
	ListCategoryByMenu(ctx context.Context, menuID string) ([]*models.Category, error)
	AddItem(ctx context.Context, categoryID, name string) (*models.Item, error)
	UpdateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, itemID string) error
	ListItemByCategory(ctx context.Context, categoryID string) ([]*models.Item, error)
	ListItemByIDs(ctx context.Context, itemIds []string) ([]*models.Item, error)
}

type menuSvcImpl struct {
	menuRepo repositories.MenuRepository
}

func NewMenuService(menuRepo repositories.MenuRepository) MenuService {
	return &menuSvcImpl{
		menuRepo: menuRepo,
	}
}

func (s *menuSvcImpl) CreateMenu(ctx context.Context, entityID, name string) (*models.Menu, error) {

	menu := models.Menu{
		ID:        rand.UUID(),
		EntityID:  entityID,
		Name:      name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.menuRepo.SaveMenu(ctx, &menu)

	if err != nil {
		return nil, err
	}
	return &menu, nil
}
func (s *menuSvcImpl) UpdateMenu(ctx context.Context, menu_ID, name string) error {

	menu, err := s.menuRepo.GetMenu(ctx, menu_ID)

	if err != nil {
		return err
	}

	menu.Name = name
	menu.UpdatedAt = time.Now().UTC()

	return s.menuRepo.UpdateMenu(ctx, menu)
}
func (s *menuSvcImpl) DeleteMenu(ctx context.Context, menuID string) error {

	return s.menuRepo.DeleteMenu(ctx, menuID)
}

func (s *menuSvcImpl) GetMenu(ctx context.Context, menuID string) (*models.Menu, error) {

	return s.menuRepo.GetMenu(ctx, menuID)
}
func (s *menuSvcImpl) GetMenuByEntity(ctx context.Context, entityID string) (*models.Menu, error) {

	return s.menuRepo.GetMenuByEntity(ctx, entityID)
}
func (s *menuSvcImpl) AddCategory(ctx context.Context, menuID, name string) (*models.Category, error) {

	category := models.Category{
		ID:          rand.UUID(),
		MenuID:      menuID,
		Name:        name,
		Description: "",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err := s.menuRepo.SaveCategory(ctx, &category)

	if err != nil {
		return nil, err
	}

	return &category, nil
}
func (s *menuSvcImpl) UpdateCategory(ctx context.Context, categoryID, name string) error {

	category, err := s.menuRepo.GetCategory(ctx, categoryID)

	if err != nil {
		return err
	}
	category.Name = name
	category.UpdatedAt = time.Now().UTC()

	return s.menuRepo.UpdateCategory(ctx, category)
}
func (s *menuSvcImpl) DeleteCategory(ctx context.Context, categoryID string) error {

	return s.menuRepo.DeleteCategory(ctx, categoryID)
}
func (s *menuSvcImpl) ListCategoryByMenu(ctx context.Context, menuID string) ([]*models.Category, error) {

	menues, err := s.menuRepo.ListCategoryByMenu(ctx, menuID)

	if err != nil {
		return nil, err
	}
	return menues, nil
}
func (s *menuSvcImpl) AddItem(ctx context.Context, categoryID, name string) (*models.Item, error) {

	item := models.Item{
		ID:          rand.UUID(),
		CategoryID:  categoryID,
		Name:        name,
		Description: "",
		ImageUrl:    "",
		Price:       0,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err := s.menuRepo.SaveItem(ctx, &item)

	if err != nil {
		return nil, err
	}
	return &item, nil
}
func (s *menuSvcImpl) UpdateItem(ctx context.Context, item *models.Item) error {

	err := s.menuRepo.UpdateItem(ctx, item)

	return err
}
func (s *menuSvcImpl) DeleteItem(ctx context.Context, itemID string) error {

	return s.menuRepo.DeleteItem(ctx, itemID)
}
func (s *menuSvcImpl) ListItemByCategory(ctx context.Context, categoryID string) ([]*models.Item, error) {

	return s.menuRepo.ListItemByCategory(ctx, categoryID)
}
func (s *menuSvcImpl) ListItemByIDs(ctx context.Context, itemIds []string) ([]*models.Item, error) {

	return s.menuRepo.ListItemsByIDs(ctx, itemIds)
}
