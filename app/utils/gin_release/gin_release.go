package gin_release

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/variable"
	"Y-frame/app/utils/response"
	"errors"
	"fmt"
	"io/ioutil"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// ReleaseRouter 根据 gin 路由包官方的建议，gin 路由引擎如果在生产模式使用，官方建议设置为 release 模式
// 官方原版提示说明：[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
// 这里我们将按照官方指导进行生产模式精细化处理
func ReleaseRouter() *gin.Engine {
	//切换成 release 模式
	gin.SetMode(gin.ReleaseMode)
	//禁用 gin 输出接口输出日志。
	gin.DefaultWriter = ioutil.Discard
	//创建路由对象
	engine := gin.New()
	//自定义中间件，对可能发生的错误进行拦截统一记录。
	engine.Use(gin.Logger(), CustomRecovery())
	return engine
}

// CustomRecovery 自定义错误(panic等)拦截中间件、对可能发生的错误进行拦截、统一记录
func CustomRecovery() gin.HandlerFunc {
	DefaultErrorWriter := &PanicExceptionRecord{}
	//对发生的错误进行恢复，并记录到日志中
	return gin.RecoveryWithWriter(DefaultErrorWriter, func(c *gin.Context, err interface{}) {
		// 这里针对发生的panic等异常进行统一响应即可
		// 这里的 err 数据类型为 ：runtime.boundsError  ，需要转为普通数据类型才可以输出
		response.ErrorsSystem(c, fmt.Sprintf("%s", err))
	})
}

//PanicExceptionRecord  panic等异常记录
type PanicExceptionRecord struct{}

func (p *PanicExceptionRecord) Write(b []byte) (n int, err error) {
	errStr := string(b)
	err = errors.New(errStr)
	//记录错误
	variable.ZapLog.Error(consts.ServerOccurredErrorMsg, zap.String("msg", errStr))
	return len(errStr), err
}
