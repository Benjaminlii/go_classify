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
