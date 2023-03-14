package util

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Benjaminlii/go_classify/biz/constants"
	"github.com/Benjaminlii/go_classify/biz/constants/errors"
	"github.com/Benjaminlii/go_classify/biz/domain/model"
	"github.com/Benjaminlii/go_classify/biz/drivers"
	uuid "github.com/satori/go.uuid"
)

// AddUserToken 向redis中添加某个用户的token，有效时间为3天
func AddUserToken(user *model.User) (token string) {
	// 生成该用户的token
	token = uuid.NewV4().String()
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Printf("[system][redis] json marshal error, err:%s", err)
		panic(errors.JSON_ERROR)
	}
	drivers.RedisClient.Set(constants.REDIS_USER_TOKEN_PRE+token, userJson, time.Hour*24*3)
	return token
}

// AddManagerToken 向redis中添加某个管理员的token，有效时间为30分钟
func AddManagerToken(administrator *model.Administrator) (token string) {
	// 生成该管理员的token
	token = uuid.NewV4().String()
	userJson, err := json.Marshal(administrator)
	if err != nil {
		log.Printf("[system][redis] json marshal error, err:%s", err)
		panic(errors.JSON_ERROR)
	}
	drivers.RedisClient.Set(constants.REDIS_MANAGER_TOKEN_PRE+token, userJson, time.Minute*30)
	return token
}
