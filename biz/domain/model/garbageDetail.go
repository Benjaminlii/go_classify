package model

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/domain/model/base"
)

// 垃圾详情
type GarbageDetail struct {
	gorm.Model
	base.Row
	Name       string `gorm:"column:name;unique_index"`     // 名称
	Path       string `gorm:"column:path"`                  // 所在类目的完整路径
	BaseTypeId uint   `gorm:"column:base_type"`             // 所属基类类型
	ImageId    uint   `gorm:"column:image_id"`              // 图片id
	Content    string `gorm:"column:content;type:longtext"` // 类目详细介绍
	Process    string `gorm:"column:process;type:longtext"` // 该类目垃圾的处理方式
}
