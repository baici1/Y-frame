package files

import (
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"mime/multipart"
	"net/http"
)

// 返回值说明：
//	7z、exe、doc 类型会返回 application/octet-stream  未知的文件类型
//	jpg	=>	image/jpeg
//	png	=>	image/png
//	ico	=>	image/x-icon
//	bmp	=>	image/bmp
//  xlsx、docx 、zip	=>	application/zip
//  tar.gz	=>	application/x-gzip
//  txt、json、log等文本文件	=>	text/plain; charset=utf-8   备注：就算txt是gbk、ansi编码，也会识别为utf-8

// GetFilesMimeByFp 这里放置一些对于 file 的一些工具函数
func GetFilesMimeByFp(fp multipart.File) string {
	//保存文件的前32位
	buffer := make([]byte, 32)
	//通过读取文件前32位获取类型
	if _, err := fp.Read(buffer); err != nil {
		variable.ZapLog.Error(g_errors.ErrorsFilesUploadReadFail + err.Error())
		return ""
	}
	return http.DetectContentType(buffer)
}
