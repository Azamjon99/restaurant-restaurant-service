package utils

import (
	"github.com/Azamjon99/restaurant-restaurant-service/src/infrastructure/repositories/sort"

	"gorm.io/gorm"
)

func Sort(sortMethod string) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		orderBy := "created_at DESC"

		if sortMethod == sort.ByCreatedAtAsc {
			orderBy = "created_at ASC"
		}
		return db.Order(orderBy)
	}
}
