package gorm_v2

import (
	"Y-frame/app/global/variable"
	"context"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

//自定义日志格式, 对 gorm 自带日志进行拦截重写
func createCustomGormLog(sqlType string, options ...Options) gormLog.Interface {
	//默认格式输出
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)
	//logger中的gormLog.Config的配置
	logConf := gormLog.Config{
		// 慢 SQL 阈值
		SlowThreshold: time.Second * variable.Configs.Gormv2.Mysql.SlowThreshold,
		// 禁用彩色打印
		Colorful: false,
		// 忽略ErrRecordNotFound（记录未找到）错误
		IgnoreRecordNotFoundError: true,
		// 日志级别
		LogLevel: gormLog.Warn,
	}
	//自定义的gorm日志对象
	log := &logger{
		Writer:       logOutPut{},
		Config:       logConf,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
	//对日志格式进行格式化输出
	for _, val := range options {
		val.apply(log)
	}
	return log
}

//实现gormLog.Writer的接口
type logOutPut struct{}

//拦截日志，将日志写入
func (l logOutPut) Printf(strFormat string, args ...interface{}) {
	logRes := fmt.Sprintf(strFormat, args...)
	logFlag := "gorm_v2 日志:"
	detailFlag := "详情："
	//通过匹配合成的日志字符串的前缀，然后选择性的写入日志文件。
	if strings.HasPrefix(strFormat, "[info]") || strings.HasPrefix(strFormat, "[traceStr]") {
		variable.ZapLog.Info(logFlag, zap.String(detailFlag, logRes))
	} else if strings.HasPrefix(strFormat, "[error]") || strings.HasPrefix(strFormat, "[traceErr]") {
		variable.ZapLog.Error(logFlag, zap.String(detailFlag, logRes))
	} else if strings.HasPrefix(strFormat, "[warn]") || strings.HasPrefix(strFormat, "[traceWarn]") {
		variable.ZapLog.Warn(logFlag, zap.String(detailFlag, logRes))
	}
}

// 尝试从外部重写内部相关的格式化变量
//定义操作接口
type Options interface {
	apply(*logger)
}

//定义执行函数的类
//这里是一个类（结构体）而不是类型
type OptionFunc struct {
	f func(log *logger)
}

//创建结构体
func NewOptionFunc(f func(log *logger)) *OptionFunc {
	return &OptionFunc{
		f: f,
	}
}

//接口的实现函数 ---执行类中匿名函数
func (o OptionFunc) apply(log *logger) {
	o.f(log)
}

// 定义 6 个函数修改内部变量
func SetInfoStrFormat(format string) Options {
	return NewOptionFunc(
		func(log *logger) {
			log.infoStr = format
		})
}

func SetWarnStrFormat(format string) Options {
	return NewOptionFunc(func(log *logger) {
		log.warnStr = format
	})
}

func SetErrStrFormat(format string) Options {
	return NewOptionFunc(func(log *logger) {
		log.errStr = format
	})
}

func SetTraceStrFormat(format string) Options {
	return NewOptionFunc(func(log *logger) {
		log.traceStr = format
	})
}
func SetTracWarnStrFormat(format string) Options {
	return NewOptionFunc(func(log *logger) {
		log.traceWarnStr = format
	})
}

func SetTracErrStrFormat(format string) Options {
	return NewOptionFunc(func(log *logger) {
		log.traceErrStr = format
	})
}

//Logger 需要实现以下接口，它接受 context，所以你可以用它来追踪日志
type logger struct {
	gormLog.Writer
	gormLog.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (l *logger) LogMode(level gormLog.LogLevel) gormLog.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l logger) Info(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Info {
		l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l logger) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Warn {
		l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l logger) Error(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Error {
		l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l logger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= gormLog.Error:
			sql, rows := fc()
			if rows == -1 {
				l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-1", sql)
			} else {
				l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormLog.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-1", sql)
			} else {
				l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.LogLevel >= gormLog.Info:
			sql, rows := fc()
			if rows == -1 {
				l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-1", sql)
			} else {
				l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}
