package dao

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/domain/model"
)

// GetRecordByUserIdLimit 获取某用户id对应的用户的识别记录，根据index和limit分页
func GetRecordByUserIdLimit(userId uint, index uint, count uint) []model.ClassifyRecord {
	db := GetDB()
	db = filterByUserId(db, userId)
	db = limit(db, index, count)
	return findClassifyRecord(db)
}

// findOrder 根据传入的db查询order
func findClassifyRecord(db *gorm.DB) (ans []model.ClassifyRecord) {
	db.Find(&ans)
	return
}
