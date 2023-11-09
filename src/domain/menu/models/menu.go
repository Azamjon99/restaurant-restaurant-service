package models

import "time"

type Menu struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	EntityID  string    `json:"entity_id" gorm:"index:idx_entity_menu"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
