package web

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/variable"
	"Y-frame/app/service/upload_files"
	"Y-frame/app/utils/response"

	"github.com/gin-gonic/gin"
)

type Upload struct {
}

//文件上传模块

func (u Upload) UploadAFile(ctx *gin.Context) {
	//保存文件路径
	savePath := variable.BasePath + variable.Configs.File.UploadFileSavePath
	if flag, finalSavePath := upload_files.UploadAFile(ctx, savePath); flag {
		//这里需要根据需求存储存储文件的路径到数据库中
		response.Success(ctx, consts.CurdStatusOkMsg, gin.H{
			"path": finalSavePath,
		})
	} else {
		response.Fail(ctx, consts.FilesUploadFailCode, consts.FilesUploadFailMsg)
	}
}
