package dao

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/domain/model"
	"log"
)

// GetRecyclesByUserIdLimit 根据用户id分页查询
func GetRecyclesByUserIdLimit(userId uint, index uint, count uint) []model.Recycle {
	db := GetDB()
	db = filterByUserId(db, userId)
	db = limit(db, index, count)
	return findRecycle(db)
}

// InsertRecycle 插入一个recycle对象
func InsertRecycle(insertRecycle *model.Recycle) *model.Recycle {
	db := GetDB()
	db.Create(insertRecycle)
	if err := db.Error; err != nil {
		log.Printf("[service][recycle][insertRecycle] db insert error, err:%s", err)
		panic(err)
	}
	return insertRecycle
}

// findRecycle 根据传入的db查询recycle
func findRecycle(db *gorm.DB) (ans []model.Recycle) {
	db.Find(&ans)
	return
}
