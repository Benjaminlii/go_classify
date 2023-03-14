package service

import (
	"github.com/Benjaminlii/go_classify/biz/dao"
	"github.com/Benjaminlii/go_classify/biz/domain/dto"
	"github.com/Benjaminlii/go_classify/biz/domain/model"
	"github.com/gin-gonic/gin"
)

// GetChildGarbageType 获取某类目下的所有子类目部分信息
func GetChildGarbageType(c *gin.Context, parentGarbageTypeId uint) ([]dto.GetChildGarbageTypeDTO, *model.GarbageType) {
	// 获取子类目信息
	childGarbageTypes := dao.FindGarbageTypeByParentId(parentGarbageTypeId)
	childs := make([]dto.GetChildGarbageTypeDTO, 0)
	for _, childGarbageType := range childGarbageTypes {
		image := dao.GetImageById(childGarbageType.ImageId)
		getChildGarbageTypeDTO := dto.GetChildGarbageTypeDTO{
			GarbageTypeId: childGarbageType.ID,
			Name:          childGarbageType.Name,
			Image:         image.Url,
		}
		childs = append(childs, getChildGarbageTypeDTO)
	}

	// 获取父类目信息，用于返回上一层
	parent := dao.GetGarbageTypeById(parentGarbageTypeId)

	return childs, parent
}

// GetGarbageDetail 通过garbageTypeId获取其对应的GarbageDetail
func GetGarbageDetail(c *gin.Context, garbageTypeId uint) *dto.GetGarbageDetailDTO {
	// 获取garbageType
	garbageType := dao.GetGarbageTypeById(garbageTypeId)

	// 根据garbageDetailId获取detail
	garbageDetail := dao.GetGarbageDetailById(garbageType.GarbageDetailId)

	// 基础类目
	baseGarbageType := dao.GetGarbageTypeById(garbageDetail.BaseTypeId)

	// 类目介绍图片
	image := dao.GetImageById(garbageDetail.ImageId)

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
