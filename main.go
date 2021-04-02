package main

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/config"
	"go_classify/biz/drivers"
)

// 主函数
func main() {
	r := gin.Default()

	config.InitConfig("conf/config.yml")
	drivers.InitFromConfigOnce()

	register(r)
	r.Run(":9090")
}