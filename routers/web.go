package routers

import (
	"Y-frame/app/http/validator/core/factory"

	"github.com/gin-gonic/gin"
)

func InitWebRouter() *gin.Engine {
	//var router *gin.Engine
	router := gin.Default()
	router.POST("/login", factory.Create("Login"))
	return router
}
