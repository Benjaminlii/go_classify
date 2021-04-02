package drivers

import (
	"go_classify/biz/config"
	"sync"
)

var(
	once sync.Once
)

func InitFromConfigOnce() {
	once.Do(func() {
		InitMySQL(&config.AppConfig)
		InitRedis(&config.AppConfig)
	})
}

