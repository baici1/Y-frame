package model

import (
	"Y-frame/app/global/variable"
	"log"

	"go.uber.org/zap"
)

func CreateUserFactory(sqlType string) *UsersModel {
	return &UsersModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type UsersModel struct {
	BaseModel   `json:"-"`
	UserName    string `gorm:"column:user_name" json:"user_name"`
	Pass        string `json:"-"`
	Phone       string `json:"phone"`
	RealName    string `gorm:"column:real_name" json:"real_name"`
	Status      int    `json:"status"`
	Token       string `json:"token"`
	LastLoginIp string `gorm:"column:last_login_ip" json:"last_login_ip"`
}

// 用户登录,
func (u *UsersModel) Login(userName string, pass string) *UsersModel {
	sql := "select id, user_name,real_name,pass,phone  from tb_users where  user_name=?  limit 1"
	result := u.Raw(sql, userName).First(u)
	log.Println(u)
	if result.Error != nil {
		// 账号密码验证失败
		variable.ZapLog.Error("根据账号查询单条记录出错:", zap.Error(result.Error))
	} else {
		return u
	}
	return nil
}
