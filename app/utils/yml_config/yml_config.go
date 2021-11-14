package yml_config

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"Y-frame/app/utils/yml_config/ymlconfig_interf"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var lastChangeTime time.Time

func init() {
	lastChangeTime = time.Now()
}

//创建yaml配置工厂
func CreateYamlFactory(fileName ...string) ymlconfig_interf.YmlConfigInterf {
	yamlConfig := viper.New()
	//配置文件目录
	yamlConfig.AddConfigPath(variable.BasePath + "/config")
	//设置需要读取的文件名：默认是common
	if len(fileName) == 0 {
		yamlConfig.SetConfigName("common")
	} else {
		yamlConfig.SetConfigName(fileName[0])
	}
	//设置配置文件的后缀
	yamlConfig.SetConfigType("yml")

	//读取配置文件
	if err := yamlConfig.ReadInConfig(); err != nil {
		log.Fatal(g_errors.ErrorsConfigInitFail + err.Error())
	}
	//返回全局配置对象
	return &ymlConfig{
		yamlConfig,
	}

}

//配置文件对象
type ymlConfig struct {
	Viper *viper.Viper
}

// 由于 vipver 包本身对于文件的变化事件有一个bug，相关事件会被回调两次
// 常年未彻底解决，相关的 issue 清单：https://github.com/spf13/viper/issues?q=OnConfigChange  https://github.com/spf13/viper/issues/619
// 设置一个内部全局变量，记录配置文件变化时的时间点，如果两次回调事件事件差小于1秒，我们认为是第二次回调事件，而不是人工修改配置文件
// 这样就避免了 vipver 包的这个bug
//监听文件变化
func (y *ymlConfig) ConfigFileChangeListen() {
	y.Viper.OnConfigChange(func(changeEvent fsnotify.Event) {
		//记录配置文件变化时的时间点，如果两次回调事件事件差小于1秒，
		if time.Since(lastChangeTime).Seconds() >= 1 {
			//当文件操作是write就重置时间。
			if changeEvent.Op.String() == "WRITE" {
				variable.ZapLog.Info(consts.YamlConfigChange)
				lastChangeTime = time.Now()
			}
		}
	})
	//监听文件变化
	y.Viper.WatchConfig()
}

// Get 一个原始值
func (y *ymlConfig) Get(keyName string) interface{} {
	return y.Viper.Get(keyName)
}

// GetString
func (y *ymlConfig) GetString(keyName string) string {
	return y.Viper.GetString(keyName)

}

// GetBool
func (y *ymlConfig) GetBool(keyName string) bool {
	return y.Viper.GetBool(keyName)
}

// GetInt
func (y *ymlConfig) GetInt(keyName string) int {
	return y.Viper.GetInt(keyName)
}

// GetInt32
func (y *ymlConfig) GetInt32(keyName string) int32 {
	return y.Viper.GetInt32(keyName)
}

// GetInt64
func (y *ymlConfig) GetInt64(keyName string) int64 {
	return y.Viper.GetInt64(keyName)
}

// float64
func (y *ymlConfig) GetFloat64(keyName string) float64 {
	return y.Viper.GetFloat64(keyName)
}

// GetDuration
func (y *ymlConfig) GetDuration(keyName string) time.Duration {
	return y.Viper.GetDuration(keyName)
}

// GetStringSlice
func (y *ymlConfig) GetStringSlice(keyName string) []string {
	return y.Viper.GetStringSlice(keyName)
}
