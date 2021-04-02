package service

import (
	"go_classify/biz/dal"
	"go_classify/biz/domain/model"
	"log"
)

// SelectUser 查询用户信息，用于登录
func SelectUser(username string, password string) (*model.User, error) {
	db := dal.GetDB()
	db = dal.FilterByUsernameAndPassword(db, username, password)
	user, err := dal.SelectUser(db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// SignUp 用户注册
func SignUp(username string, password string, name string, userIdentity uint, category uint) (*model.User, error) {
	db := dal.GetDB()
	// 数据库事物
	tx := db.Begin()
	defer tx.Commit()

	// user对象的构造
	user := &model.User{
		Username: username,
		Password: password,
	}
	user, err := dal.InsertUser(db, user)
	if err != nil {
		log.Print("[service][user][SignUp] InsertUser fail")
		return nil, err
	}

	return user, nil
}
