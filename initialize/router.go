package initialize

import (
	"bms-go/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	engine := gin.New()
	publicGroup := engine.Group("")
	// 无需鉴权API
	router.RouterGroupApp.InitBaseRouter(publicGroup)

	// 需要鉴权的API
	privateRouter := engine.Group("")
	// todo: Use middleware
	{
		// 用户
		router.RouterGroupApp.UserRouter.InitUserRouter(privateRouter)
		// 角色
		router.RouterGroupApp.AuthorityRouter.InitAuthorityRouter(privateRouter)
		// 菜单
		router.RouterGroupApp.MenuRouter.InitMenuRouter(privateRouter)
	}

	return engine
}
