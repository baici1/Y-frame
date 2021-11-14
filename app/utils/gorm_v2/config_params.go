package gorm_v2

// 数据库参数配置，结构体
// 用于解决复杂的业务场景连接到多台服务器部署的 mysql、sqlserver、postgresql 数据库
// 具体用法参见单元测试(test/gormv2_test.go)文件，TestCustomeParamsConnMysql 函数代码段

type ConfigParams struct {
	//配置读写分离
	Write ConfigParamsDetail //写
	Read  ConfigParamsDetail //读
}
type ConfigParamsDetail struct { //配置详情
	Host     string //地址
	DataBase string //数据库名称
	Port     int    //端口
	Prefix   string //前缀
	User     string //用户名称
	Pass     string //用户密码
	Charset  string //编码格式
}
