package model

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/domain/model/base"
)

// 图片
type Image struct {
	gorm.Model
	base.Row
	url string `gorm:"column:url"` // 图片完整路径
}
