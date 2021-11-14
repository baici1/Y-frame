package bootstrap

import (
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"Y-frame/app/service/sys_log_hook"
	"Y-frame/app/utils/gorm_v2"
	"Y-frame/app/utils/yml_config"
	"Y-frame/app/utils/zap_factory"
	"log"
	"os"
)

// 检查项目必须的非编译目录是否存在，避免编译后调用的时候缺失相关目录
func checkRequiredFolders() {
	//1.检查配置文件是否存在
	if _, err := os.Stat(variable.BasePath + "/config/config.yml"); err != nil {
		log.Fatal(g_errors.ErrorsConfigYamlNotExists + err.Error())
	}
}

func init() {
	// 1. 初始化 项目根路径，参见 variable 常量包，相关路径：app\global\variable\variable.go  在包里面引入
	//2.检查配置文件以及日志目录等非编译性的必要条件
	checkRequiredFolders()
	//3.初始化表单参数验证器，注册在容器（Web、Api共用容器）
	// 4.启动针对配置文件(confgi.yml、gorm_v2.yml)变化的监听， 配置文件操作指针，初始化为全局变量
	variable.ConfigYml = yml_config.CreateYamlFactory()
	variable.ConfigYml.ConfigFileChangeListen()
	// config>gorm_v2.yml 启动文件变化监听事件
	// 5.初始化全局日志句柄，并载入日志钩子处理函数
	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)
	// 6.根据配置初始化 gorm mysql 全局 *gorm.Db
	if variable.ConfigYml.GetInt("Gormv2.Mysql.IsInitGolobalGormMysql") == 1 {
		if db, err := gorm_v2.GetOneMysqlClient(); err != nil {
			log.Fatal(g_errors.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDbMysql = db
		}
	}
	// 7.雪花算法全局变量
	// 8.websocket Hub中心启动
	// 9.casbin 依据配置文件设置参数(IsInit=1)初始化
}
