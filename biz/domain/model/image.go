package model

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/domain/model/base"
)

// 图片
type Image struct {
	gorm.Model
	base.Row
	Path string `gorm:"column:path"` // 图片在服务器的结对路径
	Url  string `gorm:"column:url"`  // 图片网络路径
}
