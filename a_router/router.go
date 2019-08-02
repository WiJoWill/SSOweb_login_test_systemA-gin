package a_router

import (
	"github.com/gin-gonic/gin"

	"web_login_test_systemA/a_controller"
)
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("a_view/*")
	router.Static("/static","./static")

	router.Use(a_controller.Validate())
	{
		router.GET("/", a_controller.AGet)
		router.POST("/", a_controller.APost)
	}
	return router;
}