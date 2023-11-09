package utils

import "gorm.io/gorm"

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}
