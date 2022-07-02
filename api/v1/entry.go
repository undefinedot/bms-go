package system

import "bms-go/service"

// ApiGroup 控制器的handler
type ApiGroup struct {
	BaseApi
}

var ApiGroupApp = new(ApiGroup) // 总路由实例

var (
	// Service层的结构体集合
	userService = service.ServiceGroupApp.UserService
)
