package system

import "bms-go/service"

// ApiGroup 控制器的handler
type ApiGroup struct {
	BaseApi
	AuthorityApi
	AuthorityMenuApi
}

var ApiGroupApp = new(ApiGroup) // 总路由实例

var (
	// Service层的结构体集合
	userService      = service.ServiceGroupApp.UserService
	authorityService = service.ServiceGroupApp.AuthorityService
	baseMenuService  = service.ServiceGroupApp.BaseMenuService
	menuService      = service.ServiceGroupApp.MenuService
)
