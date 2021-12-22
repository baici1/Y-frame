package redis_factory

import (
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"Y-frame/app/utils/yml_config"
	"Y-frame/app/utils/yml_config/ymlconfig_interf"
	"time"

	"go.uber.org/zap"

	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool
var configYml ymlconfig_interf.YmlConfigInterf

// 处于程序底层的包，init 初始化的代码段的执行会优先于上层代码，因此这里读取配置项不能使用全局配置项变量
func init() {
	configYml = yml_config.CreateYamlFactory()
	redisPool = initRedisClientPool()
}

//initRedisClientPool
/* @Description: 初始化 redis 连接池
 * @return *redis.Pool
 */
func initRedisClientPool() *redis.Pool {
	redisPool = &redis.Pool{

		MaxIdle:     configYml.GetInt("Redis.MaxIdle"),                        //最大空闲数
		MaxActive:   configYml.GetInt("Redis.MaxActive"),                      //最大活跃数
		IdleTimeout: configYml.GetDuration("Redis.IdleTimeout") * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			//此处对应redis ip及端口号
			conn, err := redis.Dial("tcp", configYml.GetString("Redis.Host")+":"+configYml.GetString("Redis.Port"))
			if err != nil {
				variable.ZapLog.Error(g_errors.ErrorsRedisInitConnFail + err.Error())
				return nil, err
			}
			//获取密码
			auth := configYml.GetString("Redis.Auth") //通过配置项设置redis密码
			if len(auth) >= 1 {
				if _, err := conn.Do("AUTH", auth); err != nil {
					_ = conn.Close()
					variable.ZapLog.Error(g_errors.ErrorsRedisAuthFail + err.Error())
				}
			}
			//开始连接
			_, _ = conn.Do("select", configYml.GetInt("Redis.IndexDb"))
			return conn, err
		},
	}
	return redisPool
}

// GetOneRedisClient 从连接池获取一个redis连接
func GetOneRedisClient() *RedisClient {
	//获取重试的次数
	maxRetryTimes := configYml.GetInt("Redis.ConnFailRetryTimes")
	var oneConn redis.Conn
	for i := 1; i <= maxRetryTimes; i++ {
		oneConn = redisPool.Get()
		if oneConn.Err() != nil {
			//当次数达到最大重试次数时候记录信息
			if i == maxRetryTimes {
				variable.ZapLog.Error(g_errors.ErrorsRedisGetConnFail, zap.Error(oneConn.Err()))
				return nil
			}
			//如果出现网络短暂的抖动，短暂休眠后，支持自动重连
			time.Sleep(time.Second * configYml.GetDuration("Redis.ReConnectInterval"))
		} else {
			//获取成功就退出
			break
		}
	}
	return &RedisClient{oneConn}
}

// RedisClient 定义一个redis客户端结构体
type RedisClient struct {
	client redis.Conn
}

// 为redis-go 客户端封装统一操作函数入口
func (r *RedisClient) Execute(cmd string, args ...interface{}) (interface{}, error) {
	return r.client.Do(cmd, args...)
}

// 释放连接到连接池
func (r *RedisClient) ReleaseOneRedisClient() {
	_ = r.client.Close()
}

//  封装几个数据类型转换的函数

//bool 类型转换
func (r *RedisClient) Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply, err)
}

//string 类型转换
func (r *RedisClient) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

//string map 类型转换
func (r *RedisClient) StringMap(reply interface{}, err error) (map[string]string, error) {
	return redis.StringMap(reply, err)
}

//strings 类型转换
func (r *RedisClient) Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

//Float64 类型转换
func (r *RedisClient) Float64(reply interface{}, err error) (float64, error) {
	return redis.Float64(reply, err)
}

//int 类型转换
func (r *RedisClient) Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

//int64 类型转换
func (r *RedisClient) Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

//int map 类型转换
func (r *RedisClient) IntMap(reply interface{}, err error) (map[string]int, error) {
	return redis.IntMap(reply, err)
}

//Int64Map 类型转换
func (r *RedisClient) Int64Map(reply interface{}, err error) (map[string]int64, error) {
	return redis.Int64Map(reply, err)
}

//int64s 类型转换
func (r *RedisClient) Int64s(reply interface{}, err error) ([]int64, error) {
	return redis.Int64s(reply, err)
}

//uint64 类型转换
func (r *RedisClient) Uint64(reply interface{}, err error) (uint64, error) {
	return redis.Uint64(reply, err)
}

//Bytes 类型转换
func (r *RedisClient) Bytes(reply interface{}, err error) ([]byte, error) {
	return redis.Bytes(reply, err)
}

// 以上封装了很多最常见类型转换函数，其他您可以参考以上格式自行封装
