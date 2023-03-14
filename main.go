package main

import (
	"github.com/Benjaminlii/go_classify/biz/config"
	"github.com/Benjaminlii/go_classify/biz/drivers"
	"github.com/Benjaminlii/go_classify/biz/middleware"
	"github.com/Benjaminlii/go_classify/biz/task"
	"github.com/gin-gonic/gin"
)

// 主函数
func main() {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	config.InitConfig("conf/config.yml")
	drivers.InitFromConfigOnce()
	task.InitTask()

	register(r)
	r.Run(":8686")
}
