package dal

import (
	gorm2 "github.com/jinzhu/gorm"
	"go_classify/biz/drivers"
)

func GetDB() *gorm2.DB {
	return drivers.DB
}