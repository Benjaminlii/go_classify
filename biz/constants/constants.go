package constants

// 系统字段
const (
	TOKEN         = "token"         // 登录令牌
	CODE          = "code"          // 标准响应中的状态码
	ERROR_MESSAGE = "error_message" // 标准响应中的错误信息
	DATA          = "data"          // 标准响应中的数据域
	CURRENT_USER  = "current_user"  // 标准响应中的数据域

)

// 业务字段
const ()

// Redis相关
const (
	REDIS_USER_TOKEN_PRE = "go_classify_user_token_"     // 当前登录的用户在redis中存储有过期时间键的key前缀
	REDIS_LOCK_KEY_PRE   = "go_classify_redis_lock_key_" // redis分布式锁key
	REDIS_DEFAULT_VALUE  = "0"                           // redis无需存储value的value值
)
