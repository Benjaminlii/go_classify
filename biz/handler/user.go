package handler

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/constants"
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
		log.Printf("[service][user][SignIn] request type error, err:%s", err)
		panic(constants.REQUEST_TYPE_ERROR)
	}

	// 校验用户信息
	user, err := service.SelectUser(username, password)
	if err != nil {
		log.Printf("[service][user][SignIn] SelectUser error, err:%s", err)
		panic(err)
	}

	// 生成并添加token到redis，存储user的json
	token, err := util.AddUserToken(user)
	if err != nil {
		log.Printf("[service][user][SignIn] add user token error, err:%s", err)
		panic(err)
	}

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
		panic(constants.REQUEST_TYPE_ERROR)
	}
	userIdentity, err := util.StringToUInt(userIdentityStr)
	if err != nil {
		log.Print("[service][user][SignUp] userIdentity is not uint")
		panic(constants.REQUEST_TYPE_ERROR)
	}
	category, err := util.StringToUInt(categoryStr)
	if err != nil {
		log.Print("[service][user][SignUp] category is not uint")
		panic(constants.REQUEST_TYPE_ERROR)
	}

	// 注册
	user, err := service.SignUp(username, password, name, userIdentity, category)
	if err != nil {
		log.Printf("[service][user][SignIn] SelectUser error, err:%s", err)
		panic(err)
	}

	// 生成并添加token到redis，存储user的json
	token, err := util.AddUserToken(user)
	if err != nil {
		log.Printf("[service][user][SignUp] add user token error, err:%s", err)
		panic(err)
	}

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
		log.Printf("[service][user][LogOut] no token")
		panic(constants.NO_TOKEN_ERROR)
	}
	deleteCount, err := drivers.RedisClient.Del(constants.REDIS_USER_TOKEN_PRE + token).Result()
	if err != nil {
		log.Printf("[service][user][LogOut] delete redis key error, err:%s", err)
		panic(err)
	}
	if deleteCount != 1 {
		log.Printf("[service][user][LogOut] delete redis count is 0")
		panic(constants.SYSTEM_ERROR)
	}
	// 设置请求响应
	respMap := map[string]interface{}{}

	c.Set(constants.DATA, respMap)
}
