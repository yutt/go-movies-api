package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yutt/go-movies-api/controller"
	"github.com/yutt/go-movies-api/middleware"
)

func ConfigureEndpoints(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		users := v1.Group("users")
		{
			users.GET("/", controller.ListUsers)
			users.GET("/:id", controller.UserDetails)
		}
		auth := v1.Group("auth")
		{
			auth.POST("/register", controller.Register)
			auth.POST("/login", controller.Login)
			auth.GET("/me", middleware.RequireLogin, controller.Me)
		}
	}
	//router.GET("/user", controller.ListUsers)
	//router.GET("user/:id", controller.UserDetails)

}
