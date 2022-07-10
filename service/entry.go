package service

type ServiceGroup struct {
	UserService
	AuthorityService
	BaseMenuService
	MenuService
}

var ServiceGroupApp = new(ServiceGroup)
