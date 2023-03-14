package drivers

import (
	"sync"

	"github.com/Benjaminlii/go_classify/biz/config"
)

var (
	once sync.Once
)

func InitFromConfigOnce() {
	once.Do(func() {
		InitMySQL(&config.AppConfig)
		InitRedis(&config.AppConfig)
	})
}
