package users

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/http/controller/web"
	"Y-frame/app/http/validator/core/data_transfer"
	"Y-frame/app/utils/response"

	"github.com/gin-gonic/gin"
)

//用户登录的参数校验
type Register struct {
	UserName string `form:"user_name" json:"user_name"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	Pass     string `form:"pass" json:"pass" binding:"required,min=6,max=20"`     //  密码为 必填，长度>=6
}

func (r Register) CheckParams(ctx *gin.Context) {
	if err := ctx.ShouldBind(&r); err != nil {
		response.ValidatorError(ctx, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(r, consts.ValidatorPrefix, ctx)
	if extraAddBindDataContext != nil {
		(&web.Users{}).Register(extraAddBindDataContext)
	} else {
		response.ErrorsSystem(ctx, "Register"+consts.ValidatorParamsToJSONFail)
	}
}
