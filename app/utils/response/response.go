package response

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/g_errors"
	"Y-frame/app/utils/validator_translation"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

//封装返回数据方式 response

//ReturnJson
/* @Description: 返回数据基本函数
 * @param ctx
 * @param httpCode http的状态码
 * @param dataCode 自定义状态码
 * @param msg 返回信息
 * @param data 可选 需要携带的数据
 */
func ReturnJson(ctx *gin.Context, httpCode int, dataCode int, msg string, data ...interface{}) {
	if len(data) > 0 {
		ctx.JSON(httpCode, gin.H{
			"code": dataCode,
			"msg":  msg,
			"data": data[0],
		})
		return
	}
	ctx.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
	})

}

//Success
/* @Description: 请求成功返回函数
 * @param ctx
 * @param msg 消息
 * @param data 数据
 */
func Success(ctx *gin.Context, msg string, data ...interface{}) {
	if len(data) > 0 {
		ReturnJson(ctx, http.StatusOK, consts.CurdStatusOkCode, msg, data[0])
	} else {
		ReturnJson(ctx, http.StatusOK, consts.CurdStatusOkCode, msg)
	}

	ctx.Abort()
}

//Fail
/* @Description:请求失败返回函数
 * @param ctx
 * @param dataCode  自定义状态码
 * @param msg 消息
 */
func Fail(ctx *gin.Context, dataCode int, msg string) {
	ReturnJson(ctx, http.StatusBadRequest, dataCode, msg)
	ctx.Abort()
}

//ValidatorError
/* @Description: validator 错误专门使用的返回器
 * @param ctx
 * @param httpCode
 * @param dataCode
 * @param msg
 * @param err
 */
func ValidatorError(ctx *gin.Context, err error) {
	// 获取validator.ValidationErrors类型的errors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		errStr := errs.Error()
		// multipart:nextpart:eof 错误表示验证器需要一些参数，但是调用者没有提交任何参数
		if strings.ReplaceAll(strings.ToLower(errStr), " ", "") == "multipart:nextpart:eof" {
			ReturnJson(ctx, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, g_errors.ErrorNotAllParamsIsBlank)
		} else {
			ReturnJson(ctx, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, errStr)
		}
	} else {
		tips := validator_translation.RemoveTopStruct(errs.Translate(validator_translation.Trans))
		ReturnJson(ctx, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, tips)
	}

	ctx.Abort()
}

//ErrorsSystem
/* @Description: 处理系统中出现的问题
 * @param c
 * @param msg
 */
func ErrorsSystem(c *gin.Context, msg string) {
	ReturnJson(c, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+msg)
	c.Abort()
}
