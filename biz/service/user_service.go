package service

import (
	"github.com/Benjaminlii/go_classify/biz/constants"
	"github.com/Benjaminlii/go_classify/biz/dao"
	"github.com/Benjaminlii/go_classify/biz/domain/model"
	"github.com/Benjaminlii/go_classify/biz/util"
	"github.com/gin-gonic/gin"
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
func SignUp(username string, password string) *model.User {
	db := dao.GetDB()
	// 数据库事物
	tx := db.Begin()
	defer tx.Commit()

	// user对象的构造
	user := &model.User{
		Username:      username,
		Password:      password,
		AvatarImageId: constants.USER_DEFAULT_AVATAR_IMAGE_ID,
	}
	user = dao.InsertUser(user)

	return user
}

// PostAvatar 提交用户头像
func PostAvatar(c *gin.Context, imagePath string, imageUrl string) *model.Image {
	// 插入image
	avatarImage := &model.Image{
		Path: imagePath,
		Url:  imageUrl,
	}
	avatarImage = dao.InsertImage(avatarImage)

	// 更新用户avatarImageId
	currentUser := util.GetCurrentUser(c)
	currentUser.AvatarImageId = avatarImage.ID
	dao.SaveUser(currentUser)

	return avatarImage
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) (string, string) {
	currentUser := util.GetCurrentUser(c)
	image := dao.GetImageById(currentUser.AvatarImageId)

	return currentUser.Username, image.Url
}

// SelectAdministrator 管理员登录
func SelectAdministrator(c *gin.Context, username string, password string) *model.Administrator {
	administrator := dao.GetAdministratorByUsernameAndPassword(username, password)
	if administrator == nil {
		return nil
	}
	return administrator
}
