package dao

import (
	"github.com/jinzhu/gorm"
	"go_classify/biz/constants/errors"
	"go_classify/biz/domain/model"
	"log"
)

// GetImageById 通过id查询image
func GetImageById(imageId uint) *model.Image {
	db := GetDB()
	db = filterById(db, imageId)
	return selectImage(db)
}

// selectImage 根据db去查询image模型
func selectImage(db *gorm.DB) *model.Image {
	image := &model.Image{}
	result := db.First(image)
	if err := result.Error; err != nil {
		log.Printf("[service][image][selectImage] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(errors.SYSTEM_ERROR)
		}
	}

	return image
}
