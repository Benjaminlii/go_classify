package dao

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/constants/errors"
	"go_classify/biz/domain/model"
	"log"
)

// GetClassifyRecordById 根据id获取某一个classifyRecord
func GetClassifyRecordById(recordId uint) *model.ClassifyRecord {
	db := GetDB()
	db = filterById(db, recordId)
	return selectClassifyRecord(db)
}

// FindClassifyRecordByUserIdLimit 获取某用户id对应的用户的识别记录，根据index和limit分页
func FindClassifyRecordByUserIdLimit(userId uint, index uint, count uint) []model.ClassifyRecord {
	db := GetDB()
	db = filterByUserId(db, userId)
	db = limit(db, index, count)
	return findClassifyRecord(db)
}

// findClassifyRecord 根据传入的db查询classifyRecord
func findClassifyRecord(db *gorm.DB) (ans []model.ClassifyRecord) {
	db.Find(&ans)
	return
}

// selectClassifyRecord 根据db去查询classifyRecord模型
func selectClassifyRecord(db *gorm.DB) *model.ClassifyRecord {
	classifyRecord := &model.ClassifyRecord{}
	result := db.First(classifyRecord)
	if err := result.Error; err != nil {
		log.Printf("[service][classifyRecord][selectClassifyRecord] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(errors.SYSTEM_ERROR)
		}
	}

	return classifyRecord
}
