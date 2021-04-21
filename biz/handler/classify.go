package handler

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/constants"
	"go_classify/biz/constants/errors"
	"go_classify/biz/service"
	"go_classify/biz/util"
	"log"
)

// GetRecords 查询当前账号的识别记录列表
func GetRecords(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][classify][GetRecords] request type error, err:%s", err)
		panic(err)
	}
	indexStr, haveIndex := param["index"]
	limitStr, haveLimit := param["limit"]
	if !(haveIndex && haveLimit) {
		log.Printf("[service][classify][GetRecords] has nil in index and limit")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	index := util.StringToUInt(indexStr)
	limit := util.StringToUInt(limitStr)

	// 得到该用户的识别记录列表
	getRecordsDTOs := service.GetRecords(c, index, limit)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap["records"] = getRecordsDTOs
	c.Set(constants.DATA, respMap)
}
