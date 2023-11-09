package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/menu/models"
	"github.com/Azamjon99/restaurant-restaurant-service/src/domain/menu/repositories"
	"gorm.io/gorm"
)

const (
	tableMenu     = "restaurant.menus"
	tableCategory = "restaurant.categories"
	tableItem     = "restaurant.items"
)

type menuRepoImpl struct {
	db *gorm.DB
}

func NewMenuRepoImpl(db *gorm.DB) repositories.MenuRepository {
	return &menuRepoImpl{
		db: db,
	}
}

func (repo *menuRepoImpl) SaveMenu(ctx context.Context, menu *models.Menu) error {

	result := repo.db.WithContext(ctx).Table(tableMenu).Create(menu)

	return result.Error
}
func (repo *menuRepoImpl) UpdateMenu(ctx context.Context, menu *models.Menu) error {

	result := repo.db.WithContext(ctx).Table(tableMenu).Save(menu)

	return result.Error
}
func (repo *menuRepoImpl) DeleteMenu(ctx context.Context, menuID string) error {

	var menu models.Menu
	result := repo.db.WithContext(ctx).Table(tableMenu).Delete(&menu, "id = ?", menuID)

	return result.Error
}
func (repo *menuRepoImpl) GetMenu(ctx context.Context, menuID string) (*models.Menu, error) {

	var menu models.Menu
	result := repo.db.WithContext(ctx).Table(tableMenu).First(&menu, "id = ?", menuID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &menu, nil

}
func (repo *menuRepoImpl) GetMenuByEntity(ctx context.Context, entityID string) (*models.Menu, error) {

	var menu models.Menu
	result := repo.db.WithContext(ctx).Table(tableMenu).First(&menu, "entity_id = ?", entityID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &menu, nil
}

// category
func (repo *menuRepoImpl) SaveCategory(ctx context.Context, category *models.Category) error {

	result := repo.db.WithContext(ctx).Table(tableCategory).Create(category)

	return result.Error
}
func (repo *menuRepoImpl) UpdateCategory(ctx context.Context, category *models.Category) error {

	result := repo.db.WithContext(ctx).Table(tableCategory).Save(category)

	return result.Error
}
func (repo *menuRepoImpl) GetCategory(ctx context.Context, categoryID string) (*models.Category, error) {

	var category models.Category
	result := repo.db.WithContext(ctx).Table(tableCategory).First(&category, "id = ?", categoryID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}
func (repo *menuRepoImpl) DeleteCategory(ctx context.Context, categoryID string) error {

	var category models.Category

	result := repo.db.WithContext(ctx).Table(tableCategory).Delete(&category, "id = ?", categoryID)

	return result.Error
}
func (repo *menuRepoImpl) ListCategoryByMenu(ctx context.Context, menuID string) ([]*models.Category, error) {

	var categories []*models.Category

	result := repo.db.WithContext(ctx).Table(tableCategory).Where("menu_id = ?", menuID).Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// item
func (repo *menuRepoImpl) SaveItem(ctx context.Context, item *models.Item) error {

	result := repo.db.WithContext(ctx).Table(tableItem).Create(item)

	return result.Error
}
func (repo *menuRepoImpl) UpdateItem(ctx context.Context, item *models.Item) error {

	result := repo.db.WithContext(ctx).Table(tableItem).Save(item)

	return result.Error
}
func (repo *menuRepoImpl) GetItem(ctx context.Context, itemID string) (*models.Item, error) {

	var item models.Item
	result := repo.db.WithContext(ctx).Table(tableItem).First(&item, "id = ?", itemID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}
func (repo *menuRepoImpl) DeleteItem(ctx context.Context, itemID string) error {

	var item models.Item

	result := repo.db.WithContext(ctx).Table(tableItem).Delete(&item, "id = ?", itemID)

	return result.Error
}
func (repo *menuRepoImpl) ListItemByCategory(ctx context.Context, categoryID string) ([]*models.Item, error) {

	var items []*models.Item

	result := repo.db.WithContext(ctx).Table(tableItem).Where("category_id = ?", categoryID).Find(&items)

	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}
func (repo *menuRepoImpl) ListItemsByIDs(ctx context.Context, itemsIds []string) ([]*models.Item, error) {

	var items []*models.Item
	result := repo.db.WithContext(ctx).Table(tableItem).Where("id IN ?", itemsIds).Find(&items)

	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}
