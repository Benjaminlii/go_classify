package model

import (
	"github.com/Benjaminlii/go_classify/biz/domain/model/base"
	"github.com/jinzhu/gorm"
)

// 用户
type SetAppend struct {
	gorm.Model
	base.Row
	RecordId  uint   `gorm:"column:record_id"`  //
	ImageId   uint   `gorm:"column:image_id"`   //
	Status    uint   `gorm:"column:status"`     //
	RightType string `gorm:"column:right_type"` //
}
