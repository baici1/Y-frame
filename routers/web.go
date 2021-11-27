package routers

import (
	"Y-frame/app/global/variable"
	"Y-frame/app/http/controller/captcha"
	"Y-frame/app/http/middleware/cors"
	"Y-frame/app/http/validator/common/register_validator"
	"Y-frame/app/http/validator/core/factory"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitWebRouter() *gin.Engine {
	var router *gin.Engine
	//判断当前程序模式处理相关日志
	if variable.ConfigYml.GetBool("AppDebug") {
		router = gin.Default()
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		//pprof.Register(router)
	} else {
		//1.将日志写入日志文件
		gin.DisableConsoleColor()
		f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		gin.DefaultWriter = io.MultiWriter(f)
		// 2.如果是有nginx前置做代理，基本不需要gin框架记录访问日志，开启下面一行代码，屏蔽上面的三行代码，性能提升 5%
		//gin.SetMode(gin.ReleaseMode)
	}
	//开启跨域
	router.Use(cors.Next())
	//测试运行是否成功
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "success")
	})

	// 创建一个验证码路由
	verifyCode := router.Group("captcha")
	{
		// 验证码业务，该业务无需专门校验参数，所以可以直接调用控制器
		verifyCode.GET("/", (&captcha.Captcha{}).GenerateId)        //  获取验证码ID
		verifyCode.GET("/:captcha_id", (&captcha.Captcha{}).GetImg) // 获取图像地址
		verifyCode.GET("/audio/:captcha_id", (&captcha.Captcha{}).GetAudio)
		verifyCode.GET("/:captcha_id/:captcha_value", (&captcha.Captcha{}).CheckCode) // 校验验证码
	}

	//创建user 登录注册路由
	admin := router.Group("/admin")
	{
		//使用验证码中间件
		//admin.Use(authorization.CheckCaptchaAuth())
		admin.POST("/login", factory.Create(register_validator.Login))
		admin.POST("/register", factory.Create(register_validator.Register))
	}

	return router
}
