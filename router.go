package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureEndpoints(router *gin.Engine) {
	router.GET("/user", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, users)
	})

}
