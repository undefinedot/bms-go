package router

// RouterGroup 路由总入口
type RouterGroup struct {
	BaseRouter
	UserRouter
	AuthorityRouter
}

var RouterGroupApp = new(RouterGroup)
