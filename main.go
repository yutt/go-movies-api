package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yutt/go-movies-api/docs"
	"github.com/yutt/go-movies-api/initializers"
)

func init() {
	//Load environment variables
	initializers.LoadEnv()
	//Connect to database
	initializers.ConnectToDb()
	//Sync database
	initializers.SyncDB()
	//Initialize loggers
	initializers.CreateLogs()

}

// @title	Go movies API
// @version 0.1.0.0
// @description	An API to create a collaborative list of films
// @contact.Name Alejandro Medina Díaz
// @contact.email alejandro.medina.diaz.dev@gmail.com
func main() {

	router := gin.Default()
	router.SetTrustedProxies(nil)
	ConfigureEndpoints(router)
	//Swagger configuration
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/docs/index.html")
	})
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + os.Getenv("APP_PORT"))
}
