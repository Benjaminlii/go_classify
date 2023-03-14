package model

import (
	"github.com/Benjaminlii/go_classify/biz/domain/model/base"
	"github.com/jinzhu/gorm"
)

type Administrator struct {
	gorm.Model
	base.Row
	Username string `gorm:"column:username;unique_index"` // 管理员用户名
	Password string `gorm:"column:password"`              // 密码
}
