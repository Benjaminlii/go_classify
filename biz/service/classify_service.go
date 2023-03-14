package service

import (
	"encoding/json"
	"log"

	"github.com/Benjaminlii/go_classify/biz/constants"
	"github.com/Benjaminlii/go_classify/biz/constants/errors"
	"github.com/Benjaminlii/go_classify/biz/dao"
	"github.com/Benjaminlii/go_classify/biz/domain/dto"
	"github.com/Benjaminlii/go_classify/biz/domain/model"
	"github.com/Benjaminlii/go_classify/biz/util"
	"github.com/gin-gonic/gin"
)

// GetRecords 查询当前账号的识别记录列表
func GetRecords(c *gin.Context, index uint, limit uint) []dto.GetRecordsDTO {
	// 获取当前用户信息
	user := util.GetCurrentUser(c)

	// 获取该用户的识别记录，分页
	records := dao.FindClassifyRecordByUserIdLimit(user.ID, index, limit)

	ans := make([]dto.GetRecordsDTO, 0)

	for _, record := range records {
		image := dao.GetImageById(record.ImageId)
		garbageType := dao.GetGarbageTypeById(record.GarbageTypeId)
		getRecordsDTO := dto.GetRecordsDTO{
			RecordId:        record.ID,
			ImageUrl:        image.Url,
			GarbageTypeName: garbageType.Name,
			ClassifyTime:    record.CreatedAt.UnixNano() / 1e6,
		}

		ans = append(ans, getRecordsDTO)
	}

	return ans
}

// GoClassify 进行识别
func GoClassify(c *gin.Context, imagePath string, imageUrl string) string {
	// 插入该图片到image表中
	classifyImage := &model.Image{
		Path: imagePath,
		Url:  imageUrl,
	}
	classifyImage = dao.InsertImage(classifyImage)

	// 获得识别结果
	url := constants.DO_CLASSIFY_SERVICE_URL
	paramMap := map[string]string{"image_path": imagePath}
	result := util.Post(url, paramMap, "")
	// 结果解析
	ansMap := make(map[string]uint)
	err := json.Unmarshal([]byte(result), &ansMap)
	if err != nil {
		log.Printf("[service][classify][GoClassify] do_classify service error, err:%s", err)
		panic(err)
	}
	garbageTypeId, isOk := ansMap["code"]
	if !isOk {
		log.Printf("[service][classify][GoClassify] classify service return not have code")
		panic(errors.OUTSIDE_ERROR)
	}
	// 校验结果
	garbageType := dao.GetGarbageTypeById(garbageTypeId)
	if garbageType == nil {
		log.Printf("[service][classify][GoClassify] classify service return code is not right")
		panic(errors.OUTSIDE_ERROR)
	}

	// 存储为record记录
	currentUser := util.GetCurrentUser(c)
	record := &model.ClassifyRecord{
		UserId:        currentUser.ID,
		ImageId:       classifyImage.ID,
		GarbageTypeId: garbageType.ID,
	}
	record = dao.InsertClassifyRecord(record)

	return util.UintToString(record.ID)
}

// GoClassifyResult 通过recordId获取其对应的分类结果信息
func GoClassifyResult(c *gin.Context, recordId uint) *dto.GetGarbageDetailDTO {
	// 获取record
	record := dao.GetClassifyRecordById(recordId)

	// 获取garbageType
	garbageType := dao.GetGarbageTypeById(record.GarbageTypeId)

	// 根据garbageDetailId获取detail
	garbageDetail := dao.GetGarbageDetailById(garbageType.GarbageDetailId)

	// 基础类目
	baseGarbageType := dao.GetGarbageTypeById(garbageDetail.BaseTypeId)

	// 类目介绍图片
	image := dao.GetImageById(record.ImageId)

	ans := &dto.GetGarbageDetailDTO{
		Name:     garbageDetail.Name,
		BaseType: baseGarbageType.Name,
		Path:     garbageDetail.Path,
		Image:    image.Url,
		Content:  garbageDetail.Content,
		Process:  garbageDetail.Process,
	}
	return ans
}

// Feedback 误差结果反馈
func Feedback(c *gin.Context, recordId uint, rightType string) {
	// 获取record
	record := dao.GetClassifyRecordById(recordId)

	setAppend := &model.SetAppend{
		RecordId:  record.ID,
		ImageId:   record.ImageId,
		Status:    constants.SET_APPEND_STATIC_PENDING,
		RightType: rightType,
	}

	dao.InsertSetAppend(setAppend)

	return
}
