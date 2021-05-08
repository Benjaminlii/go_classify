package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_classify/biz/constants"
	"go_classify/biz/constants/errors"
	"go_classify/biz/dao"
	"go_classify/biz/domain/dto"
	"go_classify/biz/domain/model"
	"go_classify/biz/util"
	"log"
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
	dao.InsertClassifyRecord(record)

	return util.UintToString(garbageType.ID)
}
