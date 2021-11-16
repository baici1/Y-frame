package data_transfer

import (
	"Y-frame/app/global/variable"
	"Y-frame/app/http/validator/core/interf"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

//json https://sanyuesha.com/2018/05/07/go-json/

/*
 * DataAddContext
 * @Description:  将验证器成员（字段）绑定到数据传输的上下文中，方便控制器进行获取
 * @param validatorInterface
 * @param extraAddDataPrefix
 * @param context
 * @return *gin.Context
 */
func DataAddContext(validatorInterface interf.ValidatorInterface, extraAddDataPrefix string, context *gin.Context) *gin.Context {
	var tempJson interface{}
	if tmpBytes, err1 := json.Marshal(validatorInterface); err1 == nil { //将接口进行序列化成json字符串
		if err2 := json.Unmarshal(tmpBytes, &tempJson); err2 == nil { //进行反序列化成json结构
			if value, ok := tempJson.(map[string]interface{}); ok { //进行断言 获取一个map key 是 string，value 是存储在 interface{} 内的。
				for k, v := range value {
					context.Set(extraAddDataPrefix+k, v) //将值绑定到context
				}
				// 此外给上下文追加三个键：created_at  、 updated_at  、 deleted_at ，实际根据需要自己选择获取相关键值
				curDateTime := time.Now().Format(variable.DateFormat)
				context.Set(extraAddDataPrefix+"created_at", curDateTime)
				context.Set(extraAddDataPrefix+"updated_at", curDateTime)
				context.Set(extraAddDataPrefix+"deleted_at", curDateTime)
				return context
			}
		}
	}
	return nil
}
