package dao

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/constants/errors"
	"go_classify/biz/domain/model"
	"log"
)

// GetGarbageTypeById 通过id查询garbageType
func GetGarbageTypeById(garbageTypeId uint) *model.GarbageType {
	db := GetDB()
	db = filterById(db, garbageTypeId)
	return selectGarbageType(db)
}

// FindGarbageTypeByParentId 通过parentId查询garbageType
func FindGarbageTypeByParentId(parentGarbageTypeId uint) []model.GarbageType {
	db := GetDB()
	db = filterByParentGarbageTypeId(db, parentGarbageTypeId)
	return findGarbageType(db)
}

// filterByParentGarbageTypeId 根据父类目id过滤garbageType
func filterByParentGarbageTypeId(db *gorm.DB, parentGarbageTypeId uint) *gorm.DB {
	db = db.Where("parent_type_id = ?", parentGarbageTypeId)
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
	return
}
