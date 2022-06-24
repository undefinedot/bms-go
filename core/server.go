package core

import (
	"bms-go/global"
	"bms-go/initialize"
	"log"
)

func RunServer() {
	// 注册路由
	router := initialize.Routers()
	// 获取路由配置
	log.Fatalln(router.Run(global.SYS_CONFIG.System.Addr))
}
