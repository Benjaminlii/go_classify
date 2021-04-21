package service

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/dao"
	"go_classify/biz/domain/dto"
	"go_classify/biz/util"
)

// GetRecords 查询当前账号的识别记录列表
func GetRecords(c *gin.Context, index uint, limit uint) []dto.GetRecordsDTO {
	// 获取当前用户信息
	user := util.GetCurrentUser(c)

	// 获取该用户的识别记录，分页
	records := dao.GetRecordByUserIdLimit(user.ID, index, limit)

	ans := make([]dto.GetRecordsDTO, len(records))

	for _, record := range records {
		image := dao.GetImageById(record.ImageId)
		garbageType := dao.GetGarbageTypeById(record.GarbageTypeId)
		getRecordsDTO := dto.GetRecordsDTO{
			RecordId:        record.ID,
			ImageUrl:        image.Url,
			GarbageTypeName: garbageType.Name,
			ClassifyTime:    record.CreatedAt.Unix(),
		}

		ans = append(ans, getRecordsDTO)
	}

	return ans
}
