package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yutt/go-movies-api/controller"
)

func ConfigureEndpoints(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		users := v1.Group("users")
		{
			users.GET("/", controller.ListUsers)
			users.GET("/:id", controller.UserDetails)
		}
	}
	//router.GET("/user", controller.ListUsers)
	//router.GET("user/:id", controller.UserDetails)

}
