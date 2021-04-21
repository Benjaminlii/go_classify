package model

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/domain/model/base"
)

// 回收项
type Recycle struct {
	gorm.Model
	base.Row
	UserId           uint `gorm:"column:user_id;index"`      // 提交回收的用户id
	ClassifyRecordId uint `gorm:"column:classify_record_id"` // 回收项对应的识别记录
	Static           int  `gorm:"column:static;index"`       // 回收项的处理状态
}
