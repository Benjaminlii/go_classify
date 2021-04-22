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

	classify := goClassify.Group("/classify")
	{
		classify.POST("/get_records", handler.GetRecords)
	}

	dict := goClassify.Group("/dict")
	{
		dict.POST("/get_child_garbage_type", handler.GetChildGarbageType)
		dict.POST("/get_garbage_detail", handler.GetGarbageDetail)
	}

	recycle := goClassify.Group("/recycle")
	{
		recycle.POST("/post_recycle", handler.PostRecycle)
		recycle.POST("/get_recycles", handler.GetRecycles)
	}

	ping := goClassify.Group("/ping")
	{
		ping.POST("/ping", handler.Ping)
	}

}
