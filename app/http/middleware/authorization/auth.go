package authorization

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/variable"
	UsersToken "Y-frame/app/service/token"
	"Y-frame/app/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type HeaderParams struct {
	Authorization string `json:"authorization"  header:"Authorization" binding:"required"`
}

//CheckTokenAuth
/* @Description: 鉴权中间件
 * @return gin.HandlerFunc
 */
func CheckTokenAuth() gin.HandlerFunc {

	return func(context *gin.Context) {
		headerParams := HeaderParams{}

		//获取头部的token参数   键名：Authorization
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			response.Fail(context, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.JwtTokenMustValidMsg)
			context.Abort()
			return
		}
		//按空格分割 （头部 token字符串）
		token := strings.Split(headerParams.Authorization, " ")
		//初步判断token格式
		if len(token) == 2 && token[0] == "Bearer" && len(token[1]) >= 20 {
			claims, code := UsersToken.CreateUserToken().IsEffect(token[1])
			if code == consts.JwtTokenOK {
				//开始解析token
				key := variable.ConfigYml.GetString("Token.BindContextKeyName")
				//将token绑定到上下文
				context.Set(key, claims)
				context.Next()
			} else {
				//无效token
				response.Fail(context, http.StatusBadRequest, consts.JwtTokenInvalid, consts.JwtTokenInvalidMsg)
			}
		} else {
			//token格式错误
			response.Fail(context, http.StatusBadRequest, consts.JwtTokenFormatErrCode, consts.JwtTokenFormatErrMsg)
		}
	}
}

//RefreshToken
/* @Description: 刷新token中间件
 * @return gin.HandlerFunc
 */
func RefreshToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerParams := HeaderParams{}

		//获取头部的token参数   键名：Authorization
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			response.Fail(context, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.JwtTokenMustValidMsg)
			context.Abort()
			return
		}
		//按空格分割 （头部 token字符串）
		token := strings.Split(headerParams.Authorization, " ")
		if len(token) == 2 && token[0] == "Bearer" && len(token[1]) >= 20 {
			//检验token是否过期
			_, code := UsersToken.CreateUserToken().IsEffect(token[1])
			if code == consts.JwtTokenExpired {
				context.Next()
			} else {
				//刷新失败
				response.Fail(context, http.StatusBadRequest, consts.JwtTokenRefreshFailCode, consts.JwtTokenRefreshFailMsg)
			}
		} else {
			//token格式错误
			response.Fail(context, http.StatusBadRequest, consts.JwtTokenFormatErrCode, consts.JwtTokenFormatErrMsg)
		}
	}
}
