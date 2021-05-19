package model

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/domain/model/base"
)

// 用户识别记录
type ClassifyRecord struct {
	gorm.Model
	base.Row
	UserId        uint `gorm:"column:user_id"`         // 用户id
	ImageId       uint `gorm:"column:image_id"`        // 图片id
	GarbageTypeId uint `gorm:"column:garbage_type_id"` // 识别出的类目id
}
