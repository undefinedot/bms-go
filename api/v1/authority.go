package system

import (
	"bms-go/model"
	"bms-go/model/common/request"
	"bms-go/model/common/response"
	resModel "bms-go/model/response"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type AuthorityApi struct{}

func (a *AuthorityApi) CreateAuthority(ctx *gin.Context) {
	var authority model.Authority
	if err := ctx.ShouldBindJSON(&authority); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	if authRet, err := authorityService.CreateAuthority(authority); err != nil {
		zap.L().Error("CreateAuthority failed", zap.Error(err))
		response.FailWithMsg("创建角色失败", ctx)
	} else {
		// todo: Casbin
		response.OkWithDetail(resModel.AuthorityRes{Authority: authRet}, "创建角色成功", ctx)
	}
}

func (a *AuthorityApi) UpdateAuthority(ctx *gin.Context) {
	var authority model.Authority
	if err := ctx.ShouldBindJSON(&authority); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	if authRet, err := authorityService.UpdateAuthority(authority); err != nil {
		zap.L().Error("UpdateAuthority failed", zap.Error(err))
		response.FailWithMsg("更新角色失败", ctx)
	} else {
		response.OkWithDetail(resModel.AuthorityRes{Authority: authRet}, "更新角色成功", ctx)
	}
}

func (a *AuthorityApi) DeleteAuthority(ctx *gin.Context) {
	var authority model.Authority
	if err := ctx.ShouldBindJSON(&authority); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	if err := authorityService.DeleteAuthority(authority); err != nil {
		zap.L().Error("DeleteAuthority failed", zap.Error(err))
		response.FailWithMsg("删除角色失败", ctx)
	} else {
		response.OkWithMsg("删除角色成功", ctx)
	}
}

// GetAuthorityList 分页获取角色列表
func (a *AuthorityApi) GetAuthorityList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if err := ctx.ShouldBindJSON(&pageInfo); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	if list, total, err := authorityService.GetAuthorityInfoList(pageInfo); err != nil {
		zap.L().Error("分页获取角色列表失败", zap.Error(err))
		response.FailWithMsg("获取失败", ctx)
	} else {
		response.OkWithDetail(resModel.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}
