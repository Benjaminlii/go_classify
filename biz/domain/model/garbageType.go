package model

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/domain/model/base"
)

// 垃圾类目
type GarbageType struct {
	gorm.Model
	base.Row
	Name            string `gorm:"column:name;unique_index"` // 名称
	ParentTypeId    uint   `gorm:"column:parent_type_id"`    // 父类目id
	ImageId         uint   `gorm:"column:image_id"`          // 图片id
	GarbageDetailId uint   `gorm:"column:garbage_detail_id"` // 垃圾详情id
}
