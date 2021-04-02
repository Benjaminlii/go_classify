package model

import (
	"go_classify/biz/domain/model/base"
	"gorm.io/gorm"
)

// 用户
type User struct {
	gorm.Model
	base.Row
	Username string `gorm:"column:username;unique_index"` // 用户手机号码
	Password string `gorm:"column:password"`              // 密码
}
