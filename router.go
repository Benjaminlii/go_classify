package main

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/handler"
	"go_classify/biz/middleware"
)

func register(r *gin.Engine) {
	goClassify := r.Group("/api/go_classify")

	// 用户模块
	user := goClassify.Group("/user")
	{
		user.POST("/sign_in", handler.SignIn)
		user.POST("/sign_up", handler.SignUp)

		user.Use(middleware.CheckUserLoginMiddleware())
		user.POST("/sign_out", handler.SignOut)
	}

	goClassify.Use(middleware.CheckUserLoginMiddleware())


	ping := goClassify.Group("/ping")
	{
		ping.POST("/ping", handler.Ping)
	}

}
