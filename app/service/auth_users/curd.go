package auth_users

import (
	"Y-frame/app/model/users"
	"Y-frame/app/utils/md5_encrypt"
)

//在这里主要是处理一些复杂逻辑。

//CreateAuthUsersFactory
/* @Description: 创建工厂
 * @return *UsersCurd
 */
func CreateAuthUsersFactory() *UsersCurd {
	return &UsersCurd{}
}

type UsersCurd struct {
}

//Register
/* @Description: 用户注册 服务
 * @receiver u
 * @param userName 用户名
 * @param pass 密码
 * @param userIp 用户地址ip
 * @return bool
 */
func (u *UsersCurd) Register(userName, pass, userIp string) bool {
	pass = md5_encrypt.Base64Md5(pass) // 预先处理密码加密，然后存储在数据库
	return users.CreateUsersDBFactory("").Register(userName, pass, userIp)
}
