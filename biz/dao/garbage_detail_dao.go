package dao

import (
	"log"

	"github.com/Benjaminlii/go_classify/biz/constants/errors"
	"github.com/Benjaminlii/go_classify/biz/domain/model"
	"github.com/jinzhu/gorm"
)

// GetGarbageDetailById 根据id获取garbageDetail
func GetGarbageDetailById(garbageDetailId uint) *model.GarbageDetail {
	db := GetDB()
	db = filterById(db, garbageDetailId)
	return selectGarbageDetail(db)
}

// selectGarbageDetail 根据db去查询garbageDetail模型
func selectGarbageDetail(db *gorm.DB) *model.GarbageDetail {
	garbageDetail := &model.GarbageDetail{}
	result := db.First(garbageDetail)
	if err := result.Error; err != nil {
		log.Printf("[service][garbageDetail][selectGarbageDetail] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(errors.SYSTEM_ERROR)
		}
	}

	return garbageDetail
}
