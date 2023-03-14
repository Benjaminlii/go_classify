package util

import (
	"log"

	"github.com/Benjaminlii/go_classify/biz/constants"
	"github.com/Benjaminlii/go_classify/biz/constants/errors"
	"github.com/Benjaminlii/go_classify/biz/domain/model"
	"github.com/gin-gonic/gin"
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
