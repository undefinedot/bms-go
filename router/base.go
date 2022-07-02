package router

import (
	system "bms-go/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (b *BaseRouter) InitBaseRouter(r *gin.RouterGroup) *gin.RouterGroup {
	baseRouter := r.Group("base")
	baseApi := system.ApiGroupApp.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}

	return baseRouter
}
