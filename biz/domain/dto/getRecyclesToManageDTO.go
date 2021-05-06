package dto

type GetRecyclesToManageDTO struct {
	RecycleId     uint   `json:"recycle_id"`     // 回收项id
	ImageUrl      string `json:"image_url"`      // 回收物图片url
	Username      string `json:"username"`       // 提交回收箱的用户的用户名
	GarbageType   string `json:"garbage_type"`   // 回收项识别得到的类目名称
	RecycleTime   int64  `json:"recycle_time"`   // 提交回收项的时间
	RecycleStatic int    `json:"recycle_static"` // 回收项的处理状态
}
