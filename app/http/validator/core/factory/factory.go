package factory

import (
	"Y-frame/app/core/container"
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"Y-frame/app/http/validator/core/interf"

	"github.com/gin-gonic/gin"
)

func Create(key string) func(context *gin.Context) {
	//从容器中根据key值取出验证器对象
	if value := container.CreateContainersFactory().Get(key); value != nil {
		//通过断言拿到对象，返回验证函数
		if val, isOk := value.(interf.ValidatorInterface); isOk {
			return val.CheckParams
		}
	}
	//打印日志
	variable.ZapLog.Error(g_errors.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}
