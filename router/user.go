package router

import (
	system "bms-go/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("user")
	baseApi := system.ApiGroupApp.BaseApi
	{
		userRouter.POST("/register", baseApi.Register)             // 注册
		userRouter.POST("/changePassword", baseApi.ChangePassword) //修改密码
		userRouter.POST("/setUserInfo", baseApi.SetUserInfo)       // 修改用户信息
		userRouter.DELETE("/deleteUser", baseApi.DeleteUser)       // 删除用户
	}
	{
		userRouter.POST("/getUserList", baseApi.GetUserList) // 分页获取用户列表
		userRouter.GET("/getUserInfo", baseApi.GetUserInfo)  // 获取用户信息
	}
}
