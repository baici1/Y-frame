package upload_files

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/variable"
	"Y-frame/app/http/controller/web"
	"Y-frame/app/utils/files"
	"Y-frame/app/utils/response"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadAFile struct {
}

// CheckParams 上传文件是一个常用模块，所以很多东西都进行一个配置
/*
type FileHeader struct {
	Filename string 名字
	Header   textproto.MIMEHeader
	Size     int64 大小

	content []byte 内容
	tmpfile string
}
*/
func (u UploadAFile) CheckParams(ctx *gin.Context) {
	tmpFile, err := ctx.FormFile(variable.Configs.File.UploadFileField) //  file 是一个文件结构体（文件对象）
	//可能是上传了空文件
	if err != nil {
		response.ValidatorError(ctx, err)
		//response.Fail(ctx, consts.FilesUploadFailCode, consts.FilesUploadFailMsg)
		return
	}
	//限制文件大小，不能超过设置的最大值，同时需要进行单位转化，从 byte 转换 M。
	if tmpFile.Size > variable.Configs.File.Size<<20 {
		response.Fail(ctx, consts.FilesUploadMoreThanMaxSizeCode, consts.FilesUploadMoreThanMaxSizeMsg+strconv.FormatInt(variable.Configs.File.Size, 10))
		return
	}
	//判断文件类型，不能出现不允许的文件 mime 类型
	var flag bool = false
	if fp, err := tmpFile.Open(); err != nil {
		response.ErrorsSystem(ctx, "")
		return
	} else {
		//获取文件类型
		mimeType := files.GetFilesMimeByFp(fp)
		//与配置文件进行比较，是否是可允许上传文件类型
		for _, value := range variable.Configs.File.AllowMimeType {
			if strings.ReplaceAll(value, " ", "") == strings.ReplaceAll(mimeType, " ", "") {
				flag = true
				break
			}
		}
		//关闭文件
		_ = fp.Close()
	}
	if flag {
		(&web.Upload{}).UploadAFile(ctx)
	} else {
		response.Fail(ctx, consts.FilesUploadMimeTypeFailCode, consts.FilesUploadMimeTypeFailMsg)
	}
}
