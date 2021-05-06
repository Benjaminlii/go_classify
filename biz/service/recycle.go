package service

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/constants"
	"go_classify/biz/dao"
	"go_classify/biz/domain/dto"
	"go_classify/biz/domain/model"
	"go_classify/biz/util"
	"log"
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

// GetRecycleToMessage 查看所有用户的回收项列表
func GetRecycleToMessage(c *gin.Context, index uint, limit uint) []dto.GetRecyclesToManageDTO {
	// 分页获取回收单
	recycles := dao.GetRecyclesByStaticLimit(constants.RECYCLE_STATIC_PENDING, index, limit)

	// 结果集
	ans := make([]dto.GetRecyclesToManageDTO, len(recycles))
	for _, recycle := range recycles {
		classifyRecord := dao.GetClassifyRecordById(recycle.ClassifyRecordId)
		image := dao.GetImageById(classifyRecord.ImageId)
		postUser := dao.GetUserById(recycle.UserId)
		garbageType := dao.GetGarbageTypeById(classifyRecord.GarbageTypeId)
		getRecyclesDTO := dto.GetRecyclesToManageDTO{
			RecycleId:     recycle.ID,
			ImageUrl:      image.Url,
			Username:      postUser.Username,
			GarbageType:   garbageType.Name,
			RecycleTime:   recycle.CreatedAt.Unix(),
			RecycleStatic: recycle.Static,
		}
		ans = append(ans, getRecyclesDTO)
	}

	return ans
}

// FollowUpOrTurnDown 处理回收项，跟进或驳回
func FollowUpOrTurnDown(c *gin.Context, recycleId uint, active uint) {
	recycle := dao.GetRecycleById(recycleId)

	// 校验回收项状态
	if recycle.Static != constants.RECYCLE_STATIC_PENDING {
		log.Printf("[service][recycle][FollowUpOrTurnDown] recycle static is not pending, recycleId:%d", recycle.ID)
		panic(constants.RECYCLE_STATIC_ERROR)
	}

	if active == 1 { // 跟进
		recycle.Static = constants.RECYCLE_STATIC_DONE
	} else if active == 0 { // 驳回
		recycle.Static = constants.RECYCLE_STATIC_CLOSE
	}

	dao.SaveRecycle(recycle)
}
