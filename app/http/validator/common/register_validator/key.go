package register_validator

import "Y-frame/app/global/consts"

//管理所有的关于表单验证器的key

const (
	Login    string = consts.ValidatorPrefix + "Login"
	Register string = consts.ValidatorPrefix + "Register"
	List     string = consts.ValidatorPrefix + "List"
)
