package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/menu/models"
)

type MenuRepository interface {
	// menu
	SaveMenu(ctx context.Context, menu *models.Menu) error
	UpdateMenu(ctx context.Context, menu *models.Menu) error
	DeleteMenu(ctx context.Context, menuID string) error
	GetMenu(ctx context.Context, menuID string) (*models.Menu, error)
	GetMenuByEntity(ctx context.Context, entityID string) (*models.Menu, error)
	// category
	SaveCategory(ctx context.Context, category *models.Category) error
	GetCategory(ctx context.Context, categoryID string) (*models.Category, error)
	UpdateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, categoryID string) error
	ListCategoryByMenu(ctx context.Context, menuID string) ([]*models.Category, error)
	// item
	SaveItem(ctx context.Context, item *models.Item) error
	UpdateItem(ctx context.Context, item *models.Item) error
	GetItem(ctx context.Context, itemID string) (*models.Item, error)
	DeleteItem(ctx context.Context, itemID string) error
	ListItemByCategory(ctx context.Context, categoryID string) ([]*models.Item, error)
	ListItemsByIDs(ctx context.Context, itemsIds []string) ([]*models.Item, error)
}
