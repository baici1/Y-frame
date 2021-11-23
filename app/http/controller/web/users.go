package web

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/model"
	"Y-frame/app/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Users struct {
}

//Login
/* @Description: 用户登录API
 * @receiver u
 * @param c
 */
func (u *Users) Login(c *gin.Context) {
	//从context获取参数
	UserName := c.GetString(consts.ValidatorPrefix + "user_name")
	pass := c.GetString(consts.ValidatorPrefix + "pass")
	//创建数据库操作对象
	userModelFact := model.CreateUserFactory("")
	//进入model层
	userModel := userModelFact.Login(UserName, pass)
	if userModel != nil {
		response.Success(c, consts.CurdStatusOkMsg, userModel)
		return
	}
	response.Fail(c, http.StatusBadRequest, consts.CurdLoginFailCode, consts.CurdLoginFailMsg)
}
