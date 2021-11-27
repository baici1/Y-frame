package web

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/variable"
	"Y-frame/app/model/users"
	"Y-frame/app/service/auth_users"
	UsersToken "Y-frame/app/service/token"
	"Y-frame/app/utils/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//获取参数字段名
const (
	userNameStr = "user_name"
	passwordStr = "pass"
)

type Users struct {
}

//Register
/* @Description: 注册用户 api
 * @receiver u
 * @param ctx
 */
func (u *Users) Register(ctx *gin.Context) {
	//获取用户名和密码
	UserName := ctx.GetString(consts.ValidatorPrefix + userNameStr)
	Pass := ctx.GetString(consts.ValidatorPrefix + passwordStr)
	//获取ip地址
	UserIp := ctx.ClientIP()
	if auth_users.CreateAuthUsersFactory().Register(UserName, Pass, UserIp) {
		response.Success(ctx, consts.CurdStatusOkMsg)
	} else {
		response.Fail(ctx, http.StatusBadRequest, consts.UserRegisterFailCode, consts.UserRegisterFailMsg)
	}
}

//Login
/* @Description: 用户登录API
 * @receiver u
 * @param c
 */
func (u *Users) Login(c *gin.Context) {
	//从context获取参数
	UserName := c.GetString(consts.ValidatorPrefix + userNameStr)
	Pass := c.GetString(consts.ValidatorPrefix + passwordStr)
	//创建数据库操作对象
	userModelFact := users.CreateUsersDBFactory("")
	//进入model层
	userModel := userModelFact.Login(UserName, Pass)
	if userModel != nil {
		expireAt := variable.ConfigYml.GetInt64("Token.JwtTokenCreatedExpireAt")
		//生成token
		userTokenFactory := UsersToken.CreateUserToken()
		if token, err := userTokenFactory.GenerateToken(userModel.UserName, userModel.Phone, userModel.Id, expireAt); err == nil {
			data := gin.H{
				"userId":     userModel.Id,
				"user_name":  UserName,
				"realName":   userModel.RealName,
				"phone":      userModel.Phone,
				"token":      token,
				"updated_at": time.Now().Format(variable.DateFormat),
			}
			response.Success(c, consts.CurdStatusOkMsg, data)
		}
		return
	}
	response.Fail(c, http.StatusBadRequest, consts.UserLoginFailCode, consts.UserLoginFailMsg)
}
