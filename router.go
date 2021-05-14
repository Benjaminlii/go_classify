package main

import (
	"github.com/gin-gonic/gin"
	"go_classify/biz/handler"
	"go_classify/biz/middleware"
)

func register(r *gin.Engine) {
	goClassify := r.Group("/api/go_classify")

	// 管理员处理回收项模块
	manage := goClassify.Group("/manage")
	{
		manage.POST("/administrator_sign_up", handler.AdministratorSignUp)
		manage.Use(middleware.CheckManagerLoginMiddleware())
		manage.POST("/get_recycles_to_manage", handler.GetRecycleToMessage)
		manage.POST("/follow_up_or_turn_down", handler.FollowUpOrTurnDown)
	}

	// 用户模块
	user := goClassify.Group("/user")
	{
		user.POST("/sign_in_by_password", handler.SignIn)
		user.POST("/sign_up", handler.SignUp)
		user.POST("/administrator_sign_up", handler.AdministratorSignUp)

		user.Use(middleware.CheckUserLoginMiddleware())
		user.POST("/sign_out", handler.SignOut)
		user.POST("/post_avatar", handler.PostAvatar)
		user.POST("/get_user_info", handler.GetUserInfo)
	}

	goClassify.Use(middleware.CheckUserLoginMiddleware())

	classify := goClassify.Group("/classify")
	{
		classify.POST("/get_records", handler.GetRecords)
		classify.POST("/go_classify", handler.GoClassify)
		classify.POST("/go_classify_result", handler.GoClassifyResult)
		classify.POST("/feedback", handler.Feedback)
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
