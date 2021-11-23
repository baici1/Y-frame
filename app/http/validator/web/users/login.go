package users

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/http/controller/web"
	"Y-frame/app/http/validator/core/data_transfer"
	"Y-frame/app/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

//用户登录的参数校验
type Login struct {
	UserName string `form:"user_name" json:"user_name"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	Pass     string `form:"pass" json:"pass" binding:"required,min=6,max=20"`     //  密码为 必填，长度>=6

}

//login 验证器
func (l Login) CheckParams(c *gin.Context) {
	//获取参数值，进行初步的验证规则
	if err := c.ShouldBind(&l); err != nil {
		response.ValidatorError(c, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, err)
		return
	}
	//将参数值值绑定到context上下文中
	extraAddBindDataContext := data_transfer.DataAddContext(l, consts.ValidatorPrefix, c)
	if extraAddBindDataContext != nil {
		(&web.Users{}).Login(extraAddBindDataContext)
	} else {
		response.ErrorsSystem(c, "login 验证器参数 json 化失败")
	}
}
