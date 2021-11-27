package register_validator

import (
	"Y-frame/app/core/container"
	"Y-frame/app/http/validator/web/users"
)

func WebRegisterValidator() {
	containers := container.CreateContainersFactory()
	//登录
	containers.Set(Login, users.Login{})
	containers.Set(Register, users.Register{})
}
