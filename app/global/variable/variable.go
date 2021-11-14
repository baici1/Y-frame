package variable

import (
	"Y-frame/app/global/g_errors"
	"Y-frame/app/utils/yml_config/ymlconfig_interf"
	"log"
	"os"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

//全局变量

var (
	BasePath string //定义项目的根目录
	//配置文件
	ConfigYml ymlconfig_interf.YmlConfigInterf //全局普通配置文件指针
	//日志文件
	ZapLog *zap.Logger //全局日志指针
	//全局数据库
	GormDbMysql *gorm.DB //Mysql
)

func init() {
	if curPth, err := os.Getwd(); err != nil {
		log.Fatal(g_errors.ErrorBasePath + err.Error())
	} else {
		BasePath = curPth
	}
}
