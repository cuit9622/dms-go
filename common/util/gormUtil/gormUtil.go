package gormUtil

import "gorm.io/gorm"

func Paginate(page int32, pageSize int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 20:
			pageSize = 20
		case pageSize <= 0:
			pageSize = 5
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
