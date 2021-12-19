package upload_files

import (
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"Y-frame/app/utils/md5_encrypt"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//UploadAFile
/* @Description: 上传单个文件的服务函数
 * @param ctx
 * @param savePath
 * @return flag
 * @return finalSavePath
 */
func UploadAFile(ctx *gin.Context, savePath string) (flag bool, finalSavePath string) {
	//获取保存文件的详细目录
	newSavePath := generateYearMonthPath(savePath)
	//  1.获取上传的文件名(参数验证器已经验证完成了第一步错误，这里简化)
	file, _ := ctx.FormFile(variable.ConfigYml.GetString("FileUploadSetting.UploadFileField"))
	//2.将文件名进行加密，保证后台存储不会发生重复
	if id := variable.SnowFlake.GetId(); id > 0 {
		//组合：雪花id + 名字
		saveFileName := fmt.Sprintf("%d%s", id, file.Filename)
		//进行MD5加密
		saveFileName = md5_encrypt.MD5(saveFileName) + path.Ext(saveFileName)
		//上传文件，返回文件的相对路径
		if saveErr := ctx.SaveUploadedFile(file, newSavePath+saveFileName); saveErr == nil {
			return true, strings.ReplaceAll(newSavePath+saveFileName, variable.BasePath, "")
		}
	} else {
		variable.ZapLog.Error("文件保存出错：" + errors.New(g_errors.ErrorsSnowflakeGetIdFail).Error())
		return
	}
	return false, ""
}

//generateYearMonthPath
/* @Description: 根据时间YY_MM创建文件夹
 * @param savePath
 * @return string
 */
func generateYearMonthPath(savePath string) string {
	//获取当前时间 格式yy_mm
	currentYearMonth := time.Now().Format("2006_01")
	newSavePath := savePath + currentYearMonth
	//判断当前路径是否存在，如果不存在就重新创建
	if _, err := os.Stat(newSavePath); err != nil {
		//创建目录
		//MkdirAll 创建一个名为 path 的目录以及任何必要的父目录，并返回 nil，否则返回错误。
		if err = os.MkdirAll(newSavePath, os.ModePerm); err != nil {
			variable.ZapLog.Error("文件上传创建目录出错" + err.Error())
			return ""
		}
	}
	return newSavePath + "/"
}
