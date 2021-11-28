package register_validator

import (
	"Y-frame/app/core/container"
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
}
