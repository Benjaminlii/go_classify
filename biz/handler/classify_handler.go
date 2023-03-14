package handler

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/Benjaminlii/go_classify/biz/constants"
	"github.com/Benjaminlii/go_classify/biz/constants/errors"
	"github.com/Benjaminlii/go_classify/biz/service"
	"github.com/Benjaminlii/go_classify/biz/util"
	"github.com/gin-gonic/gin"
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

// GoClassify 进行识别
func GoClassify(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][classify][GoClassify] request type error, err:%s", err)
		panic(err)
	}
	photo, havePhoto := param["photo"]
	photoName, havePhotoName := param["photo_name"]
	if !(havePhoto && havePhotoName) {
		log.Print("[service][classify][GoClassify] has nil in photo and photoName")
		panic(errors.REQUEST_TYPE_ERROR)
	}

	photo, err = url.QueryUnescape(photo)
	if err != nil {
		log.Print("[service][classify][GoClassify] QueryUnescape error , err:", err)
		panic(err)
	}
	//photoBytes := []byte(photo)
	// 成图片文件并把文件写入到buffer
	bytes, err := base64.StdEncoding.DecodeString(photo)
	if err != nil {
		log.Print("[service][classify][GoClassify] DecodeString error , err:", err)
		panic(err)
	}
	// buffer输出到jpg文件中
	classifyPhotoName := util.GetClassifyPhotoName(c, photoName)
	imagePath := fmt.Sprintf("%s%s", constants.IMAGE_PATH_PRE_CLASSIFY_PHOTO, classifyPhotoName)
	imageUrl := fmt.Sprintf("%s%s", constants.IMAGE_URL_PRE_CLASSIFY_PHOTO, classifyPhotoName)
	err = ioutil.WriteFile(imagePath, bytes, 0666)
	if err != nil {
		log.Printf("[service][classify][GoClassify] WriteFile error, filePath:%s, err:%s", imagePath, err)
		panic(err)
	}

	// 得到该用户的识别记录列表
	garbageType := service.GoClassify(c, imagePath, imageUrl)

	// 设置请求响应
	respMap := make(map[string]interface{}, 1)
	respMap["record_id"] = garbageType
	c.Set(constants.DATA, respMap)
}

// GoClassifyResult 获取某条记录的识别分类结果
func GoClassifyResult(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][classify][GoClassifyResult] request type error, err:%s", err)
		panic(err)
	}
	recordIdStr, haveRecordId := param["record_id"]
	if !haveRecordId {
		log.Print("[service][classify][GoClassifyResult] record_id is nil")
		panic(errors.REQUEST_TYPE_ERROR)
	}

	recordId := util.StringToUInt(recordIdStr)
	// 得到该用户的识别记录列表
	garbageDetailDTO := service.GoClassifyResult(c, recordId)

	// 设置请求响应
	c.Set(constants.DATA, garbageDetailDTO)
}

// Feedback 误差结果反馈
func Feedback(c *gin.Context) {
	defer util.SetResponse(c)

	// 解析请求参数
	param := make(map[string]string)
	err := c.BindJSON(&param)
	if err != nil {
		log.Printf("[service][classify][Feedback] request type error, err:%s", err)
		panic(err)
	}
	recordIdStr, haveRecordId := param["record_id"]
	rightType, haveRightType := param["right_type"]
	if !(haveRecordId && haveRightType) {
		log.Print("[service][classify][Feedback] has nil in record_id and right_type")
		panic(errors.REQUEST_TYPE_ERROR)
	}

	recordId := util.StringToUInt(recordIdStr)
	// 得到该用户的识别记录列表
	service.Feedback(c, recordId, rightType)

	// 设置请求响应
	respMap := make(map[string]interface{}, 0)
	c.Set(constants.DATA, respMap)
}
