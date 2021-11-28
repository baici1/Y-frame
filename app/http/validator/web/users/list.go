package users

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/http/controller/web"
	"Y-frame/app/http/validator/core/data_transfer"
	"Y-frame/app/utils/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

type List struct {
	UserName string  `form:"user_name" json:"user_name"`
	Page     float64 `form:"page" json:"page" binding:"min=1"`   // 必填，页面值>=1
	Limit    float64 `form:"limit" json:"limit" binding:"min=1"` // 必填，每页条数值>=1
}

func (l List) CheckParams(ctx *gin.Context) {
	//获取参数
	if err := ctx.ShouldBind(&l); err != nil {
		fmt.Println(err)
		response.ValidatorError(ctx, err)
		return
	}
	extraAddBindDataContext := data_transfer.DataAddContext(l, consts.ValidatorPrefix, ctx)
	if extraAddBindDataContext != nil {
		(&web.Users{}).List(extraAddBindDataContext)
	} else {
		response.ErrorsSystem(ctx, "login"+consts.ValidatorParamsToJSONFail)
	}
}
