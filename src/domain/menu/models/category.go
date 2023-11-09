package models

import "time"

type Category struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	MenuID      string    `json:"menu_id" gorm:"index:idx_menu_category"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
