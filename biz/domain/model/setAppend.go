package model

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/domain/model/base"
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
