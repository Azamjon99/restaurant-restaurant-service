package models

import "time"

type RestaurantEntity struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	LogoUrl   string    `json:"logo_url"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
