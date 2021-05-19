package dto

type GetChildGarbageTypeDTO struct {
	GarbageTypeId uint   `json:"garbage_type_id"` // 子类目id
	Name          string `json:"name"`            // 子类目名
	Image         string `json:"image"`           // 子类目图片的url
}

type GetGarbageDetailDTO struct {
	Name     string `json:"name"`      // 类目的名称
	BaseType string `json:"base_type"` // 分类的基类
	Path     string `json:"path"`      // 类目的全路径
	Image    string `json:"image"`     // 图片网络url
	Content  string `json:"content"`   // 类目详细介绍
	Process  string `json:"process"`   // 该类目垃圾的处理方式
}

type GetRecordsDTO struct {
	RecordId        uint   `json:"record_id"`         // 识别记录id
	ImageUrl        string `json:"image_url"`         // 所识别图片的url
	GarbageTypeName string `json:"garbage_type_name"` // 识别出的类目名称
	ClassifyTime    int64  `json:"classify_time"`     // 进行识别的时间
}

type GetRecyclesDTO struct {
	ImageUrl      string `json:"image_url"`      // 回收的废弃物的图片url
	GarbageType   string `json:"garbage_type"`   // 废弃物识别类型
	RecycleTime   int64  `json:"recycle_time"`   // 回收单发起时间
	RecycleStatic int    `json:"recycle_static"` // 回收单发起时间
}

type GetRecyclesToManageDTO struct {
	RecycleId     uint   `json:"recycle_id"`     // 回收项id
	ImageUrl      string `json:"image_url"`      // 回收物图片url
	Username      string `json:"username"`       // 提交回收箱的用户的用户名
	GarbageType   string `json:"garbage_type"`   // 回收项识别得到的类目名称
	RecycleTime   int64  `json:"recycle_time"`   // 提交回收项的时间
	RecycleStatic int    `json:"recycle_static"` // 回收项的处理状态
}
