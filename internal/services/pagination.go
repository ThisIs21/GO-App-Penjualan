package services

import "gorm.io/gorm"

func Paginate(page, size int) func(*gorm.DB) *gorm.DB {
	if page <= 0 { page = 1 }
	if size <= 0 || size > 200 { size = 20 }
	offset := (page - 1) * size
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(size)
	}
}
