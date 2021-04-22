package handler

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/constants"
	"go_classify/biz/constants/errors"
	"go_classify/biz/service"
	"go_classify/biz/util"
	"log"
)

// PostRecycle 提交回收项
func PostRecycle(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][recycle][PostRecycle] request type error, err:%s", err)
		panic(err)
	}
	recordIdStr, haveRecordId := param["record_id"]
	if !haveRecordId {
		log.Printf("[service][recycle][PostRecycle] recordId is nil")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	recordId := util.StringToUInt(recordIdStr)

	// 得到该用户的识别记录列表
	recycleId := service.PostRecycle(c, recordId)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap["recycle_id"] = recycleId
	c.Set(constants.DATA, respMap)
}

// GetRecycles 查看用户的回收项列表
func GetRecycles(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][recycle][GetRecycles] request type error, err:%s", err)
		panic(err)
	}

	indexStr, haveIndex := param["index"]
	limitStr, haveLimit := param["limit"]
	if !(haveIndex && haveLimit) {
		log.Printf("[service][recycle][GetRecycles] has nil in index and limit")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	index := util.StringToUInt(indexStr)
	limit := util.StringToUInt(limitStr)

	// 得到该用户的识别记录列表
	ans := service.GetRecycles(c, index, limit)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap["recycles"] = ans
	c.Set(constants.DATA, respMap)
}
