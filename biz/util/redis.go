package util

import (
	"encoding/json"
	"go_classify/biz/constants"
	"go_classify/biz/constants/errors"
	"go_classify/biz/domain/model"
	"go_classify/biz/drivers"
	"log"
	"time"
)
import "github.com/satori/go.uuid"

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
