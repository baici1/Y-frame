package g_errors

const (
	//系统错误
	ErrorBasePath = "初始化项目跟目录失败"

	//容器相关
	ErrorsContainerKeyAlreadyExists string = "该键已经注册在容器中了"
	//验证器相关
	ErrorsValidatorTransInitFail string = "validator的翻译器初始化错误"
	ErrorNotAllParamsIsBlank     string = "该接口不允许所有参数都为空,请按照接口要求提交必填参数"
	ErrorsValidatorNotExists     string = "验证器不存在"
	//配置文件
	ErrorsConfigInitFail      = "初始化配置文件发生错误"
	ErrorsConfigYamlNotExists = "配置文件不存在"

	//数据库相关
	ErrorsDialectorDbInitFail      string = "gorm dialector 初始化失败,dbType:"
	ErrorsDbDriverNotExists        string = "数据库驱动类型不存在,目前支持的数据库类型：mysql，您提交数据库类型："
	ErrorsGormNotInitGlobalPointer string = "%s 数据库全局变量指针没有初始化，请在配置文件 Gormv2.yml 设置 Gormv2.%s.IsInitGolobalGormMysql = 1, 并且保证数据库配置正确 \n"
	ErrorsGormInitFail             string = "Gorm 数据库驱动、连接初始化失败"

	//token
	ErrorsTokenInvalid      string = "无效的token"
	ErrorsTokenMalFormed    string = "token格式不正确"
	ErrorsTokenNotActiveYet string = "token未激活"
	//文件上传
	ErrorsFilesUploadReadFail string = "读取文件32字节失败，详情："
	//snowflake
	ErrorsSnowflakeGetIdFail string = "获取snowflake唯一ID过程发生错误"
)
