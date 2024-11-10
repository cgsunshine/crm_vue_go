package crm

import (
	"gorm.io/gorm"
	"time"
)

// 拼接用户id 和 创建时间是搜索条件
func SearchCondition(db *gorm.DB, userId *int, startDate, endDate *time.Time) *gorm.DB {
	if startDate != nil && endDate != nil {
		db = db.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}
	if userId != nil {
		db = db.Where("user_id = ?", userId)
	}
	return db
}

func SearchConditionRole(db *gorm.DB, roleId *int, startDate, endDate *time.Time) *gorm.DB {
	if startDate != nil && endDate != nil {
		db = db.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}
	if roleId != nil {
		db = db.Where("role_id = ?", roleId)
	}
	return db
}
