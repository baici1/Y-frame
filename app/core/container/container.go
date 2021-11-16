package container

import (
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"log"
	"strings"
	"sync"
)

// 定义一个全局键值对存储容器

var sMap sync.Map

// 创建一个容器工厂
func CreateContainersFactory() *containers {
	return &containers{}
}

// 定义一个容器结构体
type containers struct {
}

//  1.以键值对的形式将代码注册到容器
func (c *containers) Set(key string, value interface{}) (res bool) {

	if _, exists := c.KeyIsExists(key); !exists {
		sMap.Store(key, value)
		res = true
	} else {
		// 程序启动阶段，zaplog 未初始化，使用系统log打印启动时候发生的异常日志
		if variable.ZapLog == nil {
			log.Fatal(g_errors.ErrorsContainerKeyAlreadyExists + ",请解决键名重复问题,相关键：" + key)
		} else {
			// 程序启动初始化完成
			variable.ZapLog.Warn(g_errors.ErrorsContainerKeyAlreadyExists + ", 相关键：" + key)
		}
	}
	return
}

//  2.删除
func (c *containers) Delete(key string) {
	sMap.Delete(key)
}

//  3.传递键，从容器获取值
func (c *containers) Get(key string) interface{} {
	if value, exists := c.KeyIsExists(key); exists {
		return value
	}
	return nil
}

//  4. 判断键是否被注册
func (c *containers) KeyIsExists(key string) (interface{}, bool) {
	return sMap.Load(key)
}

// 按照键的前缀模糊删除容器中注册的内容
func (c *containers) FuzzyDelete(keyPre string) {
	//遍历容器中的元素
	sMap.Range(func(key, value interface{}) bool {
		if keyname, ok := key.(string); ok {
			//根据前缀进行删除
			if strings.HasPrefix(keyname, keyPre) {
				sMap.Delete(keyname)
			}
		}
		return true
	})
}
