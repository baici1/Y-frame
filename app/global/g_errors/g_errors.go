package g_errors

const (
	//系统错误
	ErrorBasePath = "初始化项目跟目录失败"
	//容器相关
	ErrorsContainerKeyAlreadyExists string = "该键已经注册在容器中了"
	//验证器相关
	ErrorsValidatorNotExists string = "验证器不存在"
	//配置文件
	ErrorsConfigInitFail      = "初始化配置文件发生错误"
	ErrorsConfigYamlNotExists = "配置文件不存在"

	//数据库相关
	ErrorsDialectorDbInitFail      string = "gorm dialector 初始化失败,dbType:"
	ErrorsDbDriverNotExists        string = "数据库驱动类型不存在,目前支持的数据库类型：mysql，您提交数据库类型："
	ErrorsGormNotInitGlobalPointer string = "%s 数据库全局变量指针没有初始化，请在配置文件 Gormv2.yml 设置 Gormv2.%s.IsInitGolobalGormMysql = 1, 并且保证数据库配置正确 \n"
	ErrorsGormInitFail             string = "Gorm 数据库驱动、连接初始化失败"
)
