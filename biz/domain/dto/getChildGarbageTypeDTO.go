package dto

type GetChildGarbageTypeDTO struct {
	GarbageTypeId uint   `json:"garbage_type_id"` // 子类目id
	Name          string `json:"name"`            // 子类目名
	Image         string `json:"image"`           // 子类目图片的url
}
