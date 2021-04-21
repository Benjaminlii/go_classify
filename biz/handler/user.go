package handler

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/constants"
	"go_classify/biz/constants/errors"
	"go_classify/biz/drivers"
	"go_classify/biz/service"
	"go_classify/biz/util"
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
