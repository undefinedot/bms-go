package system

import (
	"bms-go/model"
	"bms-go/model/common/request"
	"bms-go/model/common/response"
	resModel "bms-go/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct{}

// AddBaseMenu 新增菜单
func (a *AuthorityMenuApi) AddBaseMenu(ctx *gin.Context) {
	var menu model.BaseMenu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	if err := baseMenuService.AddBaseMenu(menu); err != nil {
		zap.L().Error("AddBaseMenu() failed", zap.Error(err))
		response.FailWithMsg("添加失败", ctx)
	} else {
		response.OkWithMsg("添加成功", ctx)
	}
}

func (a *AuthorityMenuApi) DeleteBaseMenu(ctx *gin.Context) {
	var menu request.GetById
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	if err := baseMenuService.DeleteBaseMenu(menu.ID); err != nil {
		zap.L().Error("DeleteBaseMenu() failed", zap.Error(err))
		response.FailWithMsg("删除失败", ctx)
	} else {
		response.OkWithMsg("删除成功", ctx)
	}
}

func (a *AuthorityMenuApi) UpdateBaseMenu(ctx *gin.Context) {
	var menu model.BaseMenu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	if err := baseMenuService.UpdateBaseMenu(menu); err != nil {
		zap.L().Error("UpdateBaseMenu() failed", zap.Error(err))
		response.FailWithMsg("更新失败", ctx)
	} else {
		response.OkWithMsg("更新成功", ctx)
	}
}

func (a *AuthorityMenuApi) AddMenuAuthority(ctx *gin.Context) {}

func (a *AuthorityMenuApi) GetBaseMenuById(ctx *gin.Context) {
	var menuInfo request.GetById
	if err := ctx.ShouldBindJSON(&menuInfo); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	if menu, err := baseMenuService.GetBaseMenuById(menuInfo.ID); err != nil {
		zap.L().Error("GetBaseMenuById() failed", zap.Error(err))
		response.FailWithMsg("获取信息失败", ctx)
	} else {
		response.OkWithDetail(resModel.BaseMenuRes{Menu: menu}, "获取信息成功", ctx)
	}
}

// GetMenuList 分页获取 base menu 列表, 由前端来分页
func (a *AuthorityMenuApi) GetMenuList(ctx *gin.Context) {
	var pageInfo = request.PageInfo{
		Page:     1,
		PageSize: 5,
	}
	if err := ctx.ShouldBindJSON(&pageInfo); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	if menuList, total, err := menuService.GetInfoList(); err != nil {
		zap.L().Error("获取base menu的分页列表，GetInfoList() failed", zap.Error(err))
		response.FailWithMsg("获取信息失败", ctx)
	} else {
		response.OkWithDetail(resModel.PageResult{
			List:     menuList,
			Total:    total,
			PageSize: pageInfo.PageSize,
			Page:     pageInfo.Page,
		}, "获取信息成功", ctx)
	}
}

func (a *AuthorityMenuApi) GetMenuAuthority(ctx *gin.Context) {}

func (a *AuthorityMenuApi) GetMenu(ctx *gin.Context) {}

func (a *AuthorityMenuApi) GetBaseMenuTree(ctx *gin.Context) {}
