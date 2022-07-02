package main

import (
	"bms-go/core"
	"bms-go/global"
	"bms-go/initialize"
	"os"

	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	global.SYS_VP = core.Viper(os.Args...)
	// 初始化日志
	global.SYS_ZAP = core.Zap()
	zap.ReplaceGlobals(global.SYS_ZAP)
	// 初始化数据库
	global.SYS_DB = initialize.InitGorm()
	if global.SYS_DB != nil {
		initialize.RegisterTables(global.SYS_DB)
		db, _ := global.SYS_DB.DB()
		defer db.Close()
	}
	// 初始化其它
	// 注册路由
	core.RunServer()
}
