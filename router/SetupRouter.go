package router

import (
	"gintest/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter function
func SetupRouter() *gin.Engine {
	router := gin.Default()

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../view/*")
	} else {
		router.LoadHTMLGlob("view/*")
	}

	router.Static("/assetPath", "./asset")

	helloRouting := router.Group("/hello")
	{
		helloRouting.GET("", handler.GetHello)
		helloRouting.DELETE("/:id", handler.DeleteHello)
	}

	userRouting := router.Group("/user")
	{
		userRouting.GET("/:uid", handler.GetUser)
		userRouting.POST("/login", handler.UserLogin)
	}

	indexRouting := router.Group("/")
	{
		indexRouting.GET("", handler.GetIndex)
	}

	uploadSingleRouting := router.Group("/file")
	{
		uploadSingleRouting.POST("/uploadSingle", handler.UploadSingleIndex)
	}

	pingRouting := router.Group("/ping")
	{
		pingRouting.GET("", handler.WsPing)
	}

	return router
}
