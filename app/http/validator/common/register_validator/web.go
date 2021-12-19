package register_validator

import (
	"Y-frame/app/core/container"
	"Y-frame/app/http/validator/common/upload_files"
	"Y-frame/app/http/validator/web/users"
)

func WebRegisterValidator() {
	containers := container.CreateContainersFactory()
	//登录
	containers.Set(Login, users.Login{})
	//注册
	containers.Set(Register, users.Register{})
	//根据用户名查询用户信息
	containers.Set(List, users.List{})
	//上传单个文件
	containers.Set(Upload, upload_files.UploadAFile{})
}
