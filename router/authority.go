package router

import (
	system "bms-go/api/v1"

	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (ar *AuthorityRouter) InitAuthorityRouter(r *gin.RouterGroup) {
	authorityRouter := r.Group("authority")
	authorityApi := system.ApiGroupApp.AuthorityApi
	{
		authorityRouter.POST("/createAuthority", authorityApi.CreateAuthority)   // 创建角色
		authorityRouter.DELETE("/deleteAuthority", authorityApi.DeleteAuthority) // 删除角色
		authorityRouter.PUT("/updateAuthority", authorityApi.UpdateAuthority)    // 更新角色
		authorityRouter.POST("/getAuthorityList", authorityApi.GetAuthorityList) // 角色列表
	}
}
