package model

import (
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	Id        int64  `gorm:"primarykey" json:"id"`
	CreatedAt string `json:"created_at"` //日期时间字段统一设置为字符串即可
	UpdatedAt string `json:"updated_at"`
}

//创建db
func UseDbConn(sqlType string) *gorm.DB {
	var db *gorm.DB
	//获取数据库类型
	//你可以选择自己填入类型或者通过配置文件进行获取
	sqlType = strings.Trim(sqlType, " ")
	if sqlType == "" {
		sqlType = variable.Configs.Gormv2.UseDbType
	}
	switch strings.ToLower(sqlType) {
	case "mysql":
		if variable.GormDbMysql == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(g_errors.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDbMysql
	default:
		variable.ZapLog.Error(g_errors.ErrorsDbDriverNotExists + sqlType)
	}
	return db
}
