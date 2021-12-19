package routers

import (
	"Y-frame/app/global/variable"
	"Y-frame/app/http/controller/captcha"
	"Y-frame/app/http/middleware/authorization"
	"Y-frame/app/http/middleware/cors"
	"Y-frame/app/http/validator/common/register_validator"
	"Y-frame/app/http/validator/core/factory"
	"Y-frame/app/utils/gin_release"
	"net/http"

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
		//gin.DisableConsoleColor()
		//f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		//gin.DefaultWriter = io.MultiWriter(f)
		// 2.如果是有nginx前置做代理，基本不需要gin框架记录访问日志，开启下面一行代码，屏蔽上面的三行代码，性能提升 5%
		//gin.SetMode(gin.ReleaseMode)

		//【生产模式】
		// 根据 gin 官方的说明：[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
		// 如果部署到生产环境，请使用以下模式：
		// 1.生产模式(release) 和开发模式的变化主要是禁用 gin 记录接口访问日志，
		// 2.go服务就必须使用nginx作为前置代理服务，这样也方便实现负载均衡
		// 3.如果程序发生 panic 等异常使用自定义的 panic 恢复中间件拦截、记录到日志
		gin_release.ReleaseRouter()
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

		//登录注册路由
		noAuth := admin.Group("/users")
		{
			//使用验证码中间件
			//noAuth.Use(authorization.CheckCaptchaAuth())
			noAuth.POST("/login", factory.Create(register_validator.Login))
			noAuth.POST("/register", factory.Create(register_validator.Register))
		}
		//刷新token
		token := admin.Group("/token")
		{
			//先放这，到时候根据前端去考虑如何去刷新
			token.Use(authorization.CheckIsRefreshToken()).POST("/refresh")
		}
		admin.Use(authorization.CheckTokenAuth())
		{
			//用户组路由
			users := admin.Group("/users")
			{
				users.GET("/list", factory.Create(register_validator.List))
			}
		}
	}

	return router
}
