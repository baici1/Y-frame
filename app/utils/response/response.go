package response

import (
	"Y-frame/app/global/consts"
	"net/http"

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
			"data": data,
		})
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
	ReturnJson(ctx, http.StatusOK, consts.CurdStatusOkCode, msg, data)
}

//Fail
/* @Description:请求失败返回函数
 * @param ctx
 * @param httpCode  状态码
 * @param dataCode  自定义状态码
 * @param msg 消息
 */
func Fail(ctx *gin.Context, httpCode int, dataCode int, msg string) {
	ReturnJson(ctx, httpCode, dataCode, msg)
	ctx.Abort()
}
