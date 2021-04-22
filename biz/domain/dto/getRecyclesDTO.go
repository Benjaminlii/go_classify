package dto

type GetRecyclesDTO struct {
	ImageUrl      string `json:"image_url"`      // 回收的废弃物的图片url
	GarbageType   string `json:"garbage_type"`   // 废弃物识别类型
	RecycleTime   int64  `json:"recycle_time"`   // 回收单发起时间
	RecycleStatic int    `json:"recycle_static"` // 回收单发起时间
}
