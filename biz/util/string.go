package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"path"
)

// GetAvatarName 生成一个头像的name
func GetAvatarName(c *gin.Context, imageName string) string {
	return getImageName(c, imageName, "avatar")
}

// GetClassifyPhotoName 生成一个识别图片的name
func GetClassifyPhotoName(c *gin.Context, imageName string) string {
	return getImageName(c, imageName, "classify_photo")
}

// getImageName 生成一个image的name
func getImageName(c *gin.Context, imageName string, imageType string) string {
	currentUser := GetCurrentUser(c)
	// uuid截取8位
	uuidLength8 := uuid.NewV4().String()[0:8]
	// 传入文件名的后缀
	fileSuffix := path.Ext(imageName)
	newImageName := fmt.Sprintf("%s_%d_%s_%s", imageType, currentUser.ID, uuidLength8, fileSuffix)
	return newImageName
}
