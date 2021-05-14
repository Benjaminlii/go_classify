package dao

import (
	"go_classify/biz/domain/model"
	"log"
)

// InsertSetAppend 插入一个SetAppend对象
func InsertSetAppend(insertSetAppend *model.SetAppend) *model.SetAppend {
	db := GetDB()
	db = db.Create(insertSetAppend)
	if err := db.Error; err != nil {
		log.Printf("[service][user][InsertSetAppend] db insert error, err:%s", err)
		panic(err)
	}
	return insertSetAppend
}
