package main

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/config"
	"go_classify/biz/drivers"
	"go_classify/biz/task"
)

// 主函数
func main() {
	r := gin.Default()

	config.InitConfig("conf/config.yml")
	drivers.InitFromConfigOnce()
	task.InitTask()

	register(r)
	r.Run(":8686")
}
