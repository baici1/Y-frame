package casbin_v2

import (
	"Y-frame/app/global/g_errors"
	"Y-frame/app/global/variable"
	"errors"
	"fmt"
	"time"

	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/casbin/casbin/v2"
)

//创建 casbin 执行器 Enforcer
func InitCasbinEnforcer() (*casbin.SyncedEnforcer, error) {
	dsn := variable.Configs.Gormv2.Mysql.Dsn()
	fmt.Println(dsn)
	//连接数据库,获取适配器
	adapter, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/db_y_frame", true)
	if err != nil {
		return nil, errors.New(g_errors.ErrorCasbinCreateAdaptFail)
	}
	//利用适配器连接casbin
	enforcer, err := casbin.NewSyncedEnforcer(variable.Configs.Casbin.ConfPosition, adapter)
	if err != nil {
		return nil, errors.New(g_errors.ErrorCasbinCreateEnforcerFail)
	}
	// Load the policy from DB.
	_ = enforcer.LoadPolicy()
	//StartAutoLoadPolicy 启动一个 go 例程，该例程将在每个指定的持续时间调用 LoadPolicy
	enforcer.StartAutoLoadPolicy(time.Second * variable.Configs.Casbin.AutoLoadPolicySeconds)
	return enforcer, nil
}
