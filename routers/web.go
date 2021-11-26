package routers

import (
	"Y-frame/app/http/controller/captcha"
	"Y-frame/app/http/validator/common/register_validator"
	"Y-frame/app/http/validator/core/factory"

	"github.com/gin-gonic/gin"
)

func InitWebRouter() *gin.Engine {
	//var router *gin.Engine
	router := gin.Default()
	router.POST("/login", factory.Create(register_validator.Login))
	// 创建一个验证码路由
	verifyCode := router.Group("captcha")
	{
		// 验证码业务，该业务无需专门校验参数，所以可以直接调用控制器
		verifyCode.GET("/", (&captcha.Captcha{}).GenerateId)        //  获取验证码ID
		verifyCode.GET("/:captcha_id", (&captcha.Captcha{}).GetImg) // 获取图像地址
		verifyCode.GET("/audio/:captcha_id", (&captcha.Captcha{}).GetAudio)
		verifyCode.GET("/:captcha_id/:captcha_value", (&captcha.Captcha{}).CheckCode) // 校验验证码
	}
	return router
}
