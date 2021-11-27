package users

import (
	"Y-frame/app/global/variable"
	"Y-frame/app/model"
	"Y-frame/app/utils/md5_encrypt"

	"go.uber.org/zap"
)

func CreateUsersDBFactory(sqlType string) *UserModel {
	return &UserModel{BaseModel: model.BaseModel{DB: model.UseDbConn(sqlType)}}
}

type UserModel struct {
	model.BaseModel
	UserName    string `gorm:"column:user_name" json:"user_name"`
	Pass        string `json:"pass"`
	Phone       string `json:"phone"`
	RealName    string `gorm:"column:real_name" json:"real_name"`
	Status      int    `json:"status"`
	Avatar      string `gorm:"column:avatar" json:"avatar"`
	LoginTimes  int    `json:"login_times"`
	Remark      string `json:"remark"`
	LastLoginIp string `gorm:"column:last_login_ip" json:"last_login_ip"`
}

//TableName
/* @Description: 设置表名
 * @receiver u
 * @return string
 */
func (u *UserModel) TableName() string {
	return "tb_users"
}

//Register
/* @Description: 用户注册
 * @receiver u
 * @param userName 用户名
 * @param pass 密码
 * @param userIp 用户IP
 * @return bool 是否注册成功
 */
func (u *UserModel) Register(userName, pass string, userIp string) bool {
	sqlStr := "INSERT INTO tb_users(user_name,pass,last_login_ip) SELECT ?,?,? FROM DUAL WHERE NOT EXISTS(SELECT *  FROM tb_users WHERE  user_name=?)"
	result := u.Exec(sqlStr, userName, pass, userIp, userName)
	return result.RowsAffected > 0
}

//Login
/* @Description: 用户登录
 * @receiver u
 * @param userName 用户名
 * @param pass 密码
 * @return *UserModel
 */
func (u *UserModel) Login(userName string, pass string) *UserModel {
	//sql语句
	sqlStr := "select id, user_name,real_name,pass,phone  from tb_users where  user_name=?  limit 1"
	//查找满足要求的第一个
	result := u.Raw(sqlStr, userName).First(u)
	if result.Error != nil {
		// 账号验证失败
		variable.ZapLog.Error("根据账号查询单条记录出错:", zap.Error(result.Error))
		return nil
	}
	//检验密码
	if len(u.Pass) > 0 && (u.Pass == md5_encrypt.Base64Md5(pass)) {
		//记录用户登录次数
		sqlStr = "UPDATE  tb_users  set login_times=login_times+1  where  id=?"
		if result = u.Exec(sqlStr, u.Id); result.Error != nil {
			variable.ZapLog.Error("用户登录次数++出错", zap.Error(result.Error))
		}
		return u
	}
	return nil
}
