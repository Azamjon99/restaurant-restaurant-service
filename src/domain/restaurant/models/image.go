package models

import "time"

type Image struct {
	ID           string
	RestaurantID string
	Url          string
	CreatedAt    time.Time
}
