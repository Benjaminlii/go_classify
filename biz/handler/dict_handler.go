package handler

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/constants"
	"go_classify/biz/constants/errors"
	"go_classify/biz/service"
	"go_classify/biz/util"
	"log"
)

// GetChildGarbageType 查看当前类目下的子集
func GetChildGarbageType(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][dict][GetChildGarbageType] request type error, err:%s", err)
		panic(err)
	}
	parentGarbageTypeIdStr, haveParentGarbageTypeId := param["parent_garbage_type_id"]
	if !haveParentGarbageTypeId {
		log.Printf("[service][dict][GetChildGarbageType] parent garbage type id is nil")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	parentGarbageTypeId := util.StringToUInt(parentGarbageTypeIdStr)

	childGarbageTypes, parent := service.GetChildGarbageType(c, parentGarbageTypeId)

	// 设置请求响应
	respMap := make(map[string]interface{}, 2)
	respMap["child_garbage_types"] = childGarbageTypes
	respMap["parent_garbage_type_id"] = parent.ParentTypeId
	respMap["this_garbage_type_id"] = parent.ID
	c.Set(constants.DATA, respMap)
}

// GetGarbageDetail 查看当前类目下的详细类目介绍
func GetGarbageDetail(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][dict][GetGarbageDetail] request type error, err:%s", err)
		panic(err)
	}
	garbageTypeIdStr, haveGarbageTypeId := param["garbage_type_id"]
	if !haveGarbageTypeId {
		log.Printf("[service][dict][GetGarbageDetail] garbage type id is nil")
		panic(errors.REQUEST_TYPE_ERROR)
	}
	garbageTypeId := util.StringToUInt(garbageTypeIdStr)

	// 得到该用户的识别记录列表
	ans := service.GetGarbageDetail(c, garbageTypeId)

	// 设置请求响应
	c.Set(constants.DATA, ans)
}
