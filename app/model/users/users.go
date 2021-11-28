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

//getCounts
/* @Description: 根据名字查询用户表的条数 用于用户组
 * @receiver u
 * @param userName
 * @return counts
 */
func (u *UserModel) getCounts(userName string) (counts int64) {
	sqlStr := "SELECT COUNT(*) AS COUNTS FROM tb_users WHERE (user_name like ? OR real_name like ?)"
	if _ = u.Raw(sqlStr, "%"+userName+"%", "%"+userName+"%").First(&counts); counts > 0 {
		return
	}
	return 0
}

//List
/* @Description: 模糊查询 查询用户名获取用户信息
 * @receiver u
 * @param userName 用户名
 * @param limitStart  开始查询起点
 * @param limitItems 获取的条数
 * @return totalCounts 总的条数
 * @return list 用户信息
 */
func (u *UserModel) List(userName string, limitStart, limitItems int) (totalCounts int64, list []UserModel) {
	//获取用户条数
	totalCounts = u.getCounts(userName)
	if totalCounts > 0 {
		sqlStr := `SELECT  a.id, a.user_name, a.real_name,a.avatar, a.phone, a.status,a.last_login_ip,a.remark,a.login_times,
			DATE_FORMAT(created_at,'%Y-%m-%d %H:%i:%s')  AS created_at,	DATE_FORMAT(updated_at,'%Y-%m-%d %H:%i:%s')  AS updated_at  
			 FROM  tb_users a WHERE  ( user_name LIKE ? OR real_name LIKE  ?) LIMIT ?,?`
		if res := u.Raw(sqlStr, "%"+userName+"%", "%"+userName+"%", limitStart, limitItems).Find(&list); res.RowsAffected > 0 {
			return totalCounts, list
		}
		return totalCounts, nil
	}
	return 0, nil
}
