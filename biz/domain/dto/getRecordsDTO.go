package dto

type GetRecordsDTO struct {
	RecordId        uint   `json:"record_id"`         // 识别记录id
	ImageUrl        string `json:"image_url"`         // 所识别图片的url
	GarbageTypeName string `json:"garbage_type_name"` // 识别出的类目名称
	ClassifyTime    int64  `json:"classify_time"`     // 进行识别的时间
}
