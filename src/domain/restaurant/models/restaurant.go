package models

import (
	"math"
	"time"
)

type Restaurant struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	EntityID    string    `json:"entity_id" gorm:"index:idx_entity_restaurant"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     *Address  `json:"address" gorm:"serializer:json"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type Address struct {
	Name        string
	Description string
	Longitude   float64
	Latitude    float64
}

type RestaurantLocationDoc struct {
	RestaurantID string    `json:"restaurant_id"`
	Location     []float64 `json:"location"` // [lon, lat]
	EntityID     string    `json:"entity_id"`
	Name         string    `json:"name"`
	LogoUrl      string    `json:"logo_url"`
	ImageUrl     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
}
type RestaurantId struct {
	RestaurantID string `json:"restaurant_id"`
}

func (r *Restaurant) Updated() {

	if r.Address != nil {
		r.Address.Longitude = math.Round(r.Address.Longitude*math.Pow(10, float64(6))) / math.Pow(10, float64(6))
		r.Address.Latitude = math.Round(r.Address.Latitude*math.Pow(10, float64(6))) / math.Pow(10, float64(6))
	}
	r.UpdatedAt = time.Now()
}
