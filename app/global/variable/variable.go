package variable

import (
	"Y-frame/app/global/g_errors"
	"Y-frame/app/utils/snow_flake/snow_flake_interf"
	"Y-frame/app/utils/yml_config/ymlconfig_interf"
	"Y-frame/configs"
	"log"
	"os"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

//全局变量

var (
	BasePath string //定义项目的根目录
	//系统相关
	DateFormat = "2006-01-02 15:04:05" //  配置文件键值缓存时，键的前缀
	//配置文件
	ConfigYml ymlconfig_interf.YmlConfigInterf //全局普通配置文件指针
	//日志文件
	ZapLog *zap.Logger //全局日志指针
	//全局数据库
	GormDbMysql *gorm.DB //Mysql
	//雪花算法
	SnowFlake snow_flake_interf.InterfaceSnowFlake
	//全局配置变量
	Configs configs.Server
)

func init() {
	if curPth, err := os.Getwd(); err != nil {
		log.Fatal(g_errors.ErrorBasePath + err.Error())
	} else {
		BasePath = curPth
	}
}
