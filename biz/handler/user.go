package handler

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_classify/biz/constants"
	"go_classify/biz/constants/errors"
	"go_classify/biz/drivers"
	"go_classify/biz/service"
	"go_classify/biz/util"
	"io/ioutil"
	"log"
)

// SignIn 用户登录
func SignIn(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][user][SignIn] request type error, err:%s", err)
		panic(err)
	}
	username, haveUsername := param["username"]
	password, havePassword := param["password"]
	if !(haveUsername && havePassword) {
		log.Printf("[service][user][SignIn] has nil in username and password")
		panic(errors.REQUEST_TYPE_ERROR)
	}

	// 校验用户信息
	user := service.SelectUser(username, password)
	if user == nil {
		panic(errors.LOGIN_FAILD_ERROR)
	}

	// 生成并添加token到redis，存储user的json
	token := util.AddUserToken(user)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap[constants.TOKEN] = token
	c.Set(constants.DATA, respMap)
}

// SignUp 用户注册
func SignUp(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][user][SignUp] request type error, err:%s", err)
		panic(err)
	}
	username, haveUsername := param["username"]
	password, havePassword := param["password"]
	name, haveName := param["name"]
	userIdentityStr, haveUserIdentity := param["userIdentity"]
	categoryStr, haveCategory := param["category"]
	if !(haveUsername && havePassword && haveName && haveUserIdentity && haveCategory) {
		log.Print("[service][user][SignUp] has nil in username, password, name, userIdentity and category")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	userIdentity := util.StringToUInt(userIdentityStr)
	category := util.StringToUInt(categoryStr)

	// 注册
	user := service.SignUp(username, password, name, userIdentity, category)

	// 生成并添加token到redis，存储user的json
	token := util.AddUserToken(user)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap[constants.TOKEN] = token
	c.Set(constants.DATA, respMap)
}

// SignOut 登出
func SignOut(c *gin.Context) {
	defer util.SetResponse(c)

	// 从headers获取token
	token := c.Request.Header["Token"][0]
	if token == "" {
		log.Printf("[service][user][SignOut] no token")
		panic(errors.NO_TOKEN_ERROR)
	}
	deleteCount, err := drivers.RedisClient.Del(constants.REDIS_USER_TOKEN_PRE + token).Result()
	if err != nil {
		log.Printf("[service][user][LogOut] delete redis key error, err:%s", err)
		panic(err)
	}
	if deleteCount != 1 {
		log.Printf("[service][user][SignOut] delete redis count is 0")
		panic(errors.SYSTEM_ERROR)
	}
	// 设置请求响应
	respMap := map[string]interface{}{}

	c.Set(constants.DATA, respMap)
}

// PostAvatar 上传头像
func PostAvatar(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][user][PostAvatar] request type error, err:%s", err)
		panic(err)
	}
	imageBit, haveImageBit := param["image_bit"]
	imageName, haveImageName := param["image_name"]
	if !(haveImageBit && haveImageName) {
		log.Print("[service][user][PostAvatar] has nil in imageBit and imageName")
		panic(errors.REQUEST_TYPE_ERROR)
	}

	// 成图片文件并把文件写入到buffer
	bytes, _ := base64.StdEncoding.DecodeString(imageBit)
	// buffer输出到jpg文件中
	avatarName := util.GetAvatarName(c, imageName)
	imagePath := fmt.Sprintf("%s%s", constants.IMAGE_PATH_PRE_USER_AVATAR, avatarName)
	imageUrl := fmt.Sprintf("%s%s", constants.IMAGE_URL_PRE_USER_AVATAR, avatarName)
	err = ioutil.WriteFile(imagePath, bytes, 0666)
	if err != nil {
		log.Printf("[service][user][PostAvatar] WriteFile error, filePath:%s, err:%s", imagePath, err)
		panic(errors.SYSTEM_ERROR)
	}

	image := service.PostAvatar(c, imagePath, imageUrl)

	// 设置请求响应
	respMap := make(map[string]interface{}, 1)
	respMap["image_url"] = image.Url
	c.Set(constants.DATA, respMap)
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	defer util.SetResponse(c)

	userName, imageUrl := service.GetUserInfo(c)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap["user_name"] = userName
	respMap["image_url"] = imageUrl
	c.Set(constants.DATA, respMap)
}

// AdministratorSignUp 管理员登录
func AdministratorSignUp(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][user][AdministratorSignUp] request type error, err:%s", err)
		panic(err)
	}
	username, haveUsername := param["username"]
	password, havePassword := param["password"]
	if !(haveUsername && havePassword) {
		log.Printf("[service][user][AdministratorSignUp] has nil in username and password")
		panic(errors.REQUEST_TYPE_ERROR)
	}

	// 校验用户信息
	user := service.SelectAdministrator(c, username, password)
	if user == nil {
		panic(errors.LOGIN_FAILD_ERROR)
	}

	// 设置请求响应
	respMap := make(map[string]interface{})
	c.Set(constants.DATA, respMap)
}
