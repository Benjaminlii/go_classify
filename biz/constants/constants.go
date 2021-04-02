package constants

// 系统字段
const (
	TOKEN         = "Token"         // 登录令牌
	CODE          = "code"          // 标准响应中的状态码
	ERROR_MESSAGE = "error_message" // 标准响应中的错误信息
	DATA          = "data"          // 标准响应中的数据域
	CURRENT_USER  = "current_user"  // context中存储当前登录用户的key

)

// Redis相关
const (
	REDIS_USER_TOKEN_PRE = "go_classify_user_token_" // 当前登录的用户在redis中存储有过期时间键的key前缀
)
