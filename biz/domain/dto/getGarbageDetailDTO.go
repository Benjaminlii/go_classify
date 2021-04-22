package dto

type GetGarbageDetailDTO struct {
	Name     string `json:"name"`      // 类目的名称
	BaseType string `json:"base_type"` // 分类的基类
	Path     string `json:"path"`      // 类目的全路径
	Image    string `json:"image"`     // 图片网络url
	Content  string `json:"content"`   // 类目详细介绍
	Process  string `json:"process"`   // 该类目垃圾的处理方式
}
