package dao

import (
	"log"

	"github.com/Benjaminlii/go_classify/biz/constants/errors"
	"github.com/Benjaminlii/go_classify/biz/domain/model"
	"github.com/jinzhu/gorm"
)

// GetRecyclesByUserIdLimit 根据用户id分页查询
func GetRecyclesByUserIdLimit(userId uint, index uint, count uint) []model.Recycle {
	db := GetDB()
	db = filterByUserId(db, userId)
	db = limit(db, index, count)
	db = orderById(db, true)
	return findRecycle(db)
}

// GetRecyclesByStaticLimit 根据回收项状态分页查询
func GetRecyclesByStaticLimit(static int, index uint, count uint) []model.Recycle {
	db := GetDB()
	db = filterByStatic(db, static)
	db = limit(db, index, count)
	return findRecycle(db)
}

// GetRecyclesById 根据回收项状态分页查询
func GetRecycleById(recycleId uint) *model.Recycle {
	db := GetDB()
	db = filterById(db, recycleId)
	recycle := selectRecycle(db)
	return recycle
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

// filterByStatic 通过回收项状态过滤
func filterByStatic(db *gorm.DB, static int) *gorm.DB {
	db = db.Where("static = ?", static)
	return db
}

// selectRecycle 根据db去查询Recycle模型
func selectRecycle(db *gorm.DB) *model.Recycle {
	recycle := &model.Recycle{}
	result := db.First(recycle)
	if err := result.Error; err != nil {
		log.Printf("[service][user][selectUser] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(errors.SYSTEM_ERROR)
		}
	}

	return recycle
}

// SaveRecycle 更新并覆盖recycle
func SaveRecycle(recycle *model.Recycle) {
	db := GetDB()
	db.Save(recycle)
}
