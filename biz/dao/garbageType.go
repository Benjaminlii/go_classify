package dao

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/constants/errors"
	"go_classify/biz/domain/model"
	"log"
)

// GetGarbageTypeById 通过id查询garbageType
func GetGarbageTypeById(garbageTypeId uint) *model.GarbageType {
	if garbageTypeId == 0 {
		return &model.GarbageType{}
	}
	db := GetDB()
	db = filterById(db, garbageTypeId)
	return selectGarbageType(db)
}

// FindGarbageTypeByParentId 通过parentId查询garbageType
func FindGarbageTypeByParentId(parentGarbageTypeId uint) []model.GarbageType {
	db := GetDB()
	db = filterByParentGarbageTypeId(db, parentGarbageTypeId)
	db = orderById(db, true)
	return findGarbageType(db)
}

// filterByParentGarbageTypeId 根据父类目id过滤garbageType
func filterByParentGarbageTypeId(db *gorm.DB, parentGarbageTypeId uint) *gorm.DB {
	// 特化逻辑，如果传入为0，那么认定为空
	if parentGarbageTypeId == 0 {
		db = db.Where("parent_type_id is NULL")
	} else {
		db = db.Where("parent_type_id = ?", parentGarbageTypeId)
	}
	return db
}

// selectGarbageType 根据db去查询garbageType模型
func selectGarbageType(db *gorm.DB) *model.GarbageType {
	garbageType := &model.GarbageType{}
	result := db.First(garbageType)
	if err := result.Error; err != nil {
		log.Printf("[service][garbageType][selectGarbageType] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(errors.SYSTEM_ERROR)
		}
	}

	return garbageType
}

// findGarbageType 根据传入的db查询garbageType
func findGarbageType(db *gorm.DB) (ans []model.GarbageType) {
	db.Find(&ans)
	return ans
}
