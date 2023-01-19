package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type user struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

var users = []user{
	{User: "user1", Password: "12345"},
	{User: "user2", Password: "12345"},
	{User: "user3", Password: "12345"},
}

func main() {

	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, users)
	})

	router.Run(":" + os.Getenv("APP_PORT"))
}
