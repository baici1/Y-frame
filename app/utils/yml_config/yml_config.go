package yml_config

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

var lastChangeTime time.Time

func init() {
	lastChangeTime = time.Now()
}

//创建yaml配置工厂
func CreateYamlFactory(fileName ...string) *YmlConfig {
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
	return &YmlConfig{
		yamlConfig,
	}

}

//配置文件对象
type YmlConfig struct {
	Viper *viper.Viper
}

// 由于 vipver 包本身对于文件的变化事件有一个bug，相关事件会被回调两次
// 常年未彻底解决，相关的 issue 清单：https://github.com/spf13/viper/issues?q=OnConfigChange  https://github.com/spf13/viper/issues/619
// 设置一个内部全局变量，记录配置文件变化时的时间点，如果两次回调事件事件差小于1秒，我们认为是第二次回调事件，而不是人工修改配置文件
// 这样就避免了 vipver 包的这个bug
//监听文件变化
func (y *YmlConfig) ConfigFileChangeListen() {
	y.Viper.OnConfigChange(func(changeEvent fsnotify.Event) {
		//记录配置文件变化时的时间点，如果两次回调事件事件差小于1秒，
		if time.Now().Sub(lastChangeTime).Seconds() >= 1 {
			//当文件操作是write就重置时间。
			if changeEvent.Op.String() == "WRITE" {
				variable.ZapLog.Info(consts.YammlConfigChange)
				lastChangeTime = time.Now()
			}
		}
	})
	//监听文件变化
	y.Viper.WatchConfig()
}
