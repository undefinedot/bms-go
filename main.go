package main

import (
	"bms-go/core"
	"bms-go/global"
	"bms-go/initialize"
	"os"
)

func main() {
	// 初始化配置
	global.SYS_VP = core.Viper(os.Args...)
	// 初始化日志
	// 初始化数据库
	global.SYS_DB = initialize.InitGorm()
	if global.SYS_DB != nil {
		db, _ := global.SYS_DB.DB()
		defer db.Close()
	}
	// 初始化其它
	// 注册路由
}
