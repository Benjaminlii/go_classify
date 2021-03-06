package util

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/constants"
	"go_classify/biz/constants/errors"
	"go_classify/biz/domain/model"
	"log"
)

// GetCurrentUser 从context中获取登录状态
func GetCurrentUser(c *gin.Context) *model.User {
	currentUserInterFace, isOk := c.Get(constants.CURRENT_USER)
	if !isOk {
		log.Printf("[system][user][GetCurrentUser] current user is not exist")
		panic(errors.NO_LOGIN_ERROR)
	}
	currentUser, isOk := currentUserInterFace.(*model.User)
	if !isOk {
		panic(errors.SYSTEM_ERROR)
	}
	return currentUser
}
