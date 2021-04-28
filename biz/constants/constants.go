package constants

// 系统字段
const (
	TOKEN         = "token"         // 登录令牌
	CODE          = "code"          // 标准响应中的状态码
	ERROR_MESSAGE = "error_message" // 标准响应中的错误信息
	DATA          = "data"          // 标准响应中的数据域
	CURRENT_USER  = "current_user"  // 标准响应中的数据域

)

// 回收单状态
const (
	RECYCLE_STATIC_PENDING = 1 // 待处理
	RECYCLE_STATIC_IN_HAND = 2 // 处理中
	RECYCLE_STATIC_DONE    = 3 // 处理完成
	RECYCLE_STATIC_CLOSE   = 4 // 回收单关闭
)

// 图像模块
const (
	IMAGE_PATH_PRE_USER_AVATAR    = "/root/image/go_classify/avatar/"                               // 用户头像存储绝对路径前缀
	IMAGE_URL_PRE_USER_AVATAR     = "http://benjaminlee.cn/nginx/image/go_classify/avatar/"         // 用户头像静态资源url
	IMAGE_PATH_PRE_CLASSIFY_PHOTO = "/root/image/go_classify/classify_photo/"                       // 待识别图片存储绝对路径前缀
	IMAGE_URL_PRE_CLASSIFY_PHOTO  = "http://benjaminlee.cn/nginx/image/go_classify/classify_photo/" // 待识别图片静态资源url

	USER_DEFAULT_AVATAR_IMAGE_ID = 1 // 用户默认头像id

	DO_CLASSIFY_SERVICE_URL = "http://localhost:8686/classify/do_classify/" // 识别服务的API地址
)

// Redis相关
const (
	REDIS_USER_TOKEN_PRE = "go_classify_user_token_"     // 当前登录的用户在redis中存储有过期时间键的key前缀
	REDIS_LOCK_KEY_PRE   = "go_classify_redis_lock_key_" // redis分布式锁key
	REDIS_DEFAULT_VALUE  = "0"                           // redis无需存储value的value值
)
