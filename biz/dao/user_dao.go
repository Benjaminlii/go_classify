package dao

import (
	"log"

	"github.com/Benjaminlii/go_classify/biz/constants/errors"
	"github.com/Benjaminlii/go_classify/biz/domain/model"
	"github.com/jinzhu/gorm"
)

// GetUserById 根据userId获取user
func GetUserById(userId uint) *model.User {
	db := GetDB()
	db = filterById(db, userId)
	user := selectUser(db)
	return user
}

// GetUserByUsernameAndPassword 根据用户名密码查找user
func GetUserByUsernameAndPassword(username string, password string) *model.User {
	db := GetDB()
	db = filterByUsernameAndPassword(db, username, password)
	user := selectUser(db)
	return user
}

// InsertUser 插入一个user对象
func InsertUser(insertUser *model.User) *model.User {
	db := GetDB()
	db = db.Create(insertUser)
	if err := db.Error; err != nil {
		log.Printf("[service][user][InsertUser] db insert error, err:%s", err)
		panic(err)
	}
	return insertUser
}

// SaveUser 更新并覆盖user
func SaveUser(user *model.User) {
	db := GetDB()
	db.Save(user)
}

// selectUser 根据db去查询user模型
func selectUser(db *gorm.DB) *model.User {
	user := &model.User{}
	result := db.First(user)
	if err := result.Error; err != nil {
		log.Printf("[service][user][selectUser] db select error, err:%s", err)
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			panic(errors.SYSTEM_ERROR)
		}
	}

	return user
}

// filterByUsernameAndPassword 通过用户名以及密码
func filterByUsernameAndPassword(db *gorm.DB, username string, password string) *gorm.DB {
	db = db.Where("username = ?", username)
	db = db.Where("password = ?", password)
	return db
}

// filterByUserId 通过userId过滤
func filterByUserId(db *gorm.DB, userId uint) *gorm.DB {
	db = db.Where("user_id = ?", userId)
	return db
}
