package gorm_v2

import (
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
)

/*
我这里并没有配置多个数据库，我现在没有这个需求。但是可以扩展,同时你可以通过结构体器配置多个数据库
*/

//创建一个mysql客户端，这里不设置读写分离，根据场景再去加。
func GetOneMysqlClient() (*gorm.DB, error) {
	//获取数据库类型
	sqlType := variable.ConfigYml.GetString("Gormv2.UseDbType")
	////设置是否读取分离
	//readDbIsOpen:=variable.ConfigYml.Viper.GetInt("Gormv2." + sqlType + ".IsOpenReadDb")
	return GetSqlDriver(sqlType)
}

// 获取数据库驱动, 可以通过options 动态参数连接任意多个数据库
func GetSqlDriver(sqlType string) (*gorm.DB, error) {
	//定义数据库驱动指针
	var dbDialector gorm.Dialector
	//获取对应数据库驱动指针
	if val, err := getDbDialector(sqlType); err != nil {
		variable.ZapLog.Error(g_errors.ErrorsDialectorDbInitFail+sqlType, zap.Error(err))
	} else {
		dbDialector = val
	}
	//开始通过gorm驱动数据库
	gormDb, err := gorm.Open(dbDialector, &gorm.Config{
		SkipDefaultTransaction: true,                 //跳过默认事务 为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。
		PrepareStmt:            true,                 //在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
		Logger:                 redefineLog(sqlType), //拦截、接管 gorm v2 自带日志
	})
	if err != nil {
		//gorm 数据库驱动初始化失败
		return nil, err
	}
	// 查询没有数据，屏蔽 gorm v2 包中会爆出的错误
	// https://github.com/go-gorm/gorm/issues/3789  此 issue 所反映的问题就是我们本次解决掉的
	_ = gormDb.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", func(d *gorm.DB) {
		d.Statement.RaiseErrorOnNotFound = false
	})
	// 为主连接设置连接池
	if rawDb, err := gormDb.DB(); err != nil {
		return nil, err
	} else {
		//相关配置
		//连接池里面的连接最大空闲时长。
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		//连接不活动时的最大生存时间(秒)
		rawDb.SetConnMaxLifetime(variable.ConfigYml.GetDuration("Gormv2."+sqlType+".Write.SetConnMaxLifetime") * time.Second)
		//设置与数据库建立连接的最大数目。
		rawDb.SetMaxIdleConns(variable.ConfigYml.GetInt("Gormv2." + sqlType + ".Write.SetMaxIdleConns"))
		//设置连接池中的最大闲置连接数。
		rawDb.SetMaxOpenConns(variable.ConfigYml.GetInt("Gormv2." + sqlType + ".Write.SetMaxOpenConns"))
		return gormDb, nil
	}
}

//获取一个数据库方言(Dialector),通俗的说就是根据不同的连接参数，获取具体的一类数据库的连接指针
func getDbDialector(sqlType string) (gorm.Dialector, error) {
	var dbDialector gorm.Dialector
	//获取数据库驱动的dsn
	dsn := getDsn(sqlType)
	switch strings.ToLower(sqlType) {
	case "mysql":
		dbDialector = mysql.Open(dsn)
	default:
		return nil, errors.New(g_errors.ErrorsDbDriverNotExists + sqlType)
	}
	return dbDialector, nil
}

//根据配置参数生成数据库驱动 dsn
func getDsn(sqlType string) string {
	readWrite := "Write"
	//获取配置文件的数据库的参数
	Host := variable.ConfigYml.GetString("Gormv2." + sqlType + "." + readWrite + ".Host")
	DataBase := variable.ConfigYml.GetString("Gormv2." + sqlType + "." + readWrite + ".DataBase")
	Port := variable.ConfigYml.GetInt("Gormv2." + sqlType + "." + readWrite + ".Port")
	User := variable.ConfigYml.GetString("Gormv2." + sqlType + "." + readWrite + ".User")
	Pass := variable.ConfigYml.GetString("Gormv2." + sqlType + "." + readWrite + ".Pass")
	Charset := variable.ConfigYml.GetString("Gormv2." + sqlType + "." + readWrite + ".Charset")
	//根据数据库进行选择配置dsn 因为常用是mysql 所以只有一个
	switch strings.ToLower(sqlType) {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", User, Pass, Host, Port, DataBase, Charset)
	default:
		return ""
	}
}

//创建自定义日志模块，对 gorm 日志进行拦截、
func redefineLog(sqlType string) gormLog.Interface {
	isOwnLogger := variable.ConfigYml.GetBool("Gormv2." + sqlType + ".Log.IsOwnLogger")
	infoStr := variable.ConfigYml.GetString("Gormv2." + sqlType + ".Log.infoStr")
	warnStr := variable.ConfigYml.GetString("Gormv2." + sqlType + ".Log.warnStr")
	errStr := variable.ConfigYml.GetString("Gormv2." + sqlType + ".Log.errStr")
	traceStr := variable.ConfigYml.GetString("Gormv2." + sqlType + ".Log.traceStr")
	traceWarnStr := variable.ConfigYml.GetString("Gormv2." + sqlType + ".Log.traceWarnStr")
	traceErrStr := variable.ConfigYml.GetString("Gormv2." + sqlType + ".Log.traceErrStr")
	if isOwnLogger {
		return createCustomGormLog(sqlType,
			SetInfoStrFormat(infoStr),
			SetWarnStrFormat(warnStr),
			SetErrStrFormat(errStr),
			SetTraceStrFormat(traceStr),
			SetTracWarnStrFormat(traceWarnStr),
			SetTracErrStrFormat(traceErrStr))

	}
	return createCustomGormLog(sqlType)
}
