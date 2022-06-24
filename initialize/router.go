package initialize

import (
	system "bms-go/api/v1"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	engine := gin.New()
	{
		// 无需鉴权API
		engine.POST("/login", system.Login)     // 首页登录
		engine.POST("/captcha", system.Captcha) // 登录时的验证码
	}

	return engine
}
