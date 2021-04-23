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
	IMAGE_PATH_PRE_USER_AVATAR = "/root/image/go_classify/avatar/"
)

// Redis相关
const (
	REDIS_USER_TOKEN_PRE = "go_classify_user_token_"     // 当前登录的用户在redis中存储有过期时间键的key前缀
	REDIS_LOCK_KEY_PRE   = "go_classify_redis_lock_key_" // redis分布式锁key
	REDIS_DEFAULT_VALUE  = "0"                           // redis无需存储value的value值
)
