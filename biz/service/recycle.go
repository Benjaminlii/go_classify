package service

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/constants"
	"go_classify/biz/dao"
	"go_classify/biz/domain/dto"
	"go_classify/biz/domain/model"
	"go_classify/biz/util"
)

// PostRecycle 提交回收项
func PostRecycle(c *gin.Context, recordId uint) (recycleId uint) {
	// 获取当前用户信息
	currentUser := util.GetCurrentUser(c)

	// 拼装回收单信息
	newRecycle := &model.Recycle{
		UserId:           currentUser.ID,
		ClassifyRecordId: recordId,
		Static:           constants.RECYCLE_STATIC_PENDING,
	}

	insertedRecycle := dao.InsertRecycle(newRecycle)

	return insertedRecycle.ID
}

// GetRecycles 查看用户的回收项列表
func GetRecycles(c *gin.Context, index uint, limit uint) []dto.GetRecyclesDTO {
	// 获取当前用户信息
	currentUser := util.GetCurrentUser(c)

	// 分页获取回收单
	recycles := dao.GetRecyclesByUserIdLimit(currentUser.ID, index, limit)

	// 结果集
	ans := make([]dto.GetRecyclesDTO, len(recycles))
	for _, recycle := range recycles {
		classifyRecord := dao.GetClassifyRecordById(recycle.ClassifyRecordId)
		image := dao.GetImageById(classifyRecord.ImageId)
		garbageType := dao.GetGarbageTypeById(classifyRecord.GarbageTypeId)
		getRecyclesDTO := dto.GetRecyclesDTO{
			ImageUrl:      image.Url,
			GarbageType:   garbageType.Name,
			RecycleTime:   recycle.CreatedAt.Unix(),
			RecycleStatic: recycle.Static,
		}
		ans = append(ans, getRecyclesDTO)
	}

	return ans
}
