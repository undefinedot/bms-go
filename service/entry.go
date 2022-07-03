package service

type ServiceGroup struct {
	UserService
	AuthorityService
}

var ServiceGroupApp = new(ServiceGroup)
