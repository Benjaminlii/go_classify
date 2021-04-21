package service

import (
	"go_classify/biz/dao"
	"go_classify/biz/domain/model"
)

// SelectUser 查询用户信息，用于登录
func SelectUser(username string, password string) *model.User {
	user := dao.GetUserByUsernameAndPassword(username, password)
	if user == nil {
		return nil
	}
	return user
}

// SignUp 用户注册
func SignUp(username string, password string, name string, userIdentity uint, category uint) *model.User {
	db := dao.GetDB()
	// 数据库事物
	tx := db.Begin()
	defer tx.Commit()

	// user对象的构造
	user := &model.User{
		Username: username,
		Password: password,
	}
	user = dao.InsertUser(user)

	return user
}
