package models

import "time"

type Item struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	CategoryID  string    `json:"category_id" gorm:"index:idx_category_item"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
