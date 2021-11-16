package main

import (
	"Y-frame/app/global/variable"
	_ "Y-frame/bootstrap"
	"Y-frame/routers"
)

//这里存放管理系统的路由 （后台管理系统）
func main() {
	router := routers.InitWebRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Web.Port"))
}
