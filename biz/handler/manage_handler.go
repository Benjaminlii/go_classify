package handler

import (
	"log"

	"github.com/Benjaminlii/go_classify/biz/constants"
	"github.com/Benjaminlii/go_classify/biz/constants/errors"
	"github.com/Benjaminlii/go_classify/biz/service"
	"github.com/Benjaminlii/go_classify/biz/util"
	"github.com/gin-gonic/gin"
)

// AdministratorSignUp 管理员登录
func AdministratorSignUp(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][manage][AdministratorSignUp] request type error, err:%s", err)
		panic(err)
	}
	username, haveUsername := param["username"]
	password, havePassword := param["password"]
	if !(haveUsername && havePassword) {
		log.Printf("[service][manage][AdministratorSignUp] has nil in username and password")
		panic(errors.REQUEST_TYPE_ERROR)
	}

	// 校验用户信息
	administrator := service.SelectAdministrator(c, username, password)
	if administrator == nil {
		panic(errors.LOGIN_FAILD_ERROR)
	}

	// 生成并添加token到redis，存储user的json
	token := util.AddManagerToken(administrator)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap[constants.TOKEN] = token
	c.Set(constants.DATA, respMap)
}

// GetRecycleToMessage 查看所有用户的回收项列表
func GetRecycleToMessage(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][manage][GetRecycleToMessage] request type error, err:%s", err)
		panic(err)
	}
	indexStr, haveIndex := param["index"]
	limitStr, haveLimit := param["limit"]
	if !(haveIndex && haveLimit) {
		log.Printf("[service][manage][GetRecycleToMessage] has nil in index and limit")
		panic(errors.REQUEST_TYPE_ERROR)
	}

	index := util.StringToUInt(indexStr)
	limit := util.StringToUInt(limitStr)

	getRecyclesToManageDTOs := service.GetRecycleToMessage(c, index, limit)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap["recycles"] = getRecyclesToManageDTOs
	c.Set(constants.DATA, respMap)
}

// FollowUpOrTurnDown 处理回收项，跟进或驳回
func FollowUpOrTurnDown(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][manage][FollowUpOrTurnDown] request type error, err:%s", err)
		panic(err)
	}
	recycleIdStr, haveRecycleId := param["recycle_id"]
	activeStr, haveActive := param["active"]
	if !(haveRecycleId && haveActive) {
		log.Printf("[service][manage][FollowUpOrTurnDown] has nil in recycle id and active")
		panic(errors.REQUEST_TYPE_ERROR)
	}

	recycleId := util.StringToUInt(recycleIdStr)
	active := util.StringToUInt(activeStr)

	service.FollowUpOrTurnDown(c, recycleId, active)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	c.Set(constants.DATA, respMap)
}
