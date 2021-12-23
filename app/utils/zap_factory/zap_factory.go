package zap_factory

import (
	"Y-frame/app/global/variable"
	"log"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
zap是管理全局日志，不包括请求的日志，更多管理请求流程中的日志情况
*/

//创建日志对象的工厂
func CreateZapFactory(entry func(zapcore.Entry) error) *zap.Logger {
	//获取当前文件配置的模式 开发模式，生产模式
	appDebug := variable.Configs.Zaps.AppDebug

	//判断当前所处的模式
	//开发模式直接返回一个便捷的zap日志管理器地址，所有的日志打印到控制台即可
	if appDebug {
		if logger, err := zap.NewDevelopment(zap.Hooks(entry)); err != nil {
			log.Fatal("创建zap日志包失败，详情：" + err.Error())
		} else {
			return logger
		}
	}
	//生产模式  自定义日志格式
	//Encoder:编码器(如何写入日志)。
	encoderConfig := zap.NewProductionEncoderConfig()
	var encoder zapcore.Encoder
	switch variable.Configs.Zaps.TextFormat {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig) // json格式
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	}
	//WriterSyncer ：指定日志将写到哪里去。
	//写入文件位置
	fileName := variable.BasePath + variable.Configs.Zaps.GoSkeletonLogName
	//配置相关信息（日志切割归档功能）
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,                         //日志文件的位置
		MaxSize:    variable.Configs.Zaps.MaxSize,    //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: variable.Configs.Zaps.MaxBackups, //保留旧文件的最大个数
		MaxAge:     variable.Configs.Zaps.MaxAge,     //保留旧文件的最大天数
		Compress:   variable.Configs.Zaps.Compress,   //是否压缩/归档旧文件
	}
	//写入器
	writer := zapcore.AddSync(lumberJackLogger)

	//配置其余选项（时间等）
	//获取时间格式化的级别 可选项有秒 毫秒 默认为秒
	timePrecision := variable.Configs.Zaps.TimePrecision
	//对时间进行格式化
	var recordTimeFormat string
	switch timePrecision {
	case "second":
		recordTimeFormat = "2006-01-02 15:04:05"
	case "millisecond":
		recordTimeFormat = "2006-01-02 15:04:05.000"
	default:
		recordTimeFormat = "2006-01-02 15:04:05"
	}
	//修改时间编码器
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(recordTimeFormat))
	}
	//在日志文件中使用大写字母记录日志级别 （日志级别都是大写）
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 开始初始化zap日志核心参数，
	//参数一：编码器
	//参数二：写入器
	//参数三：参数级别，debug级别支持后续调用的所有函数写日志，如果是 fatal 高级别，则级别>=fatal 才可以写日志
	zapCore := zapcore.NewCore(encoder, writer, zap.InfoLevel)
	return zap.New(zapCore, zap.AddCaller(), zap.Hooks(entry))
}
