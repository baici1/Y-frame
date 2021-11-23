package routers

import (
	"Y-frame/app/http/validator/common/register_validator"
	"Y-frame/app/http/validator/core/factory"

	"github.com/gin-gonic/gin"
)

func InitWebRouter() *gin.Engine {
	//var router *gin.Engine
	router := gin.Default()
	router.POST("/login", factory.Create(register_validator.Login))
	return router
}
