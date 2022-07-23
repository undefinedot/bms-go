package service

import (
	"bms-go/global"
	"bms-go/model"
	"bms-go/model/common/request"
	"strconv"
)

type MenuService struct{}

// GetMenuAuthority 获取角色的关联菜单列表
func (m *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (menus []model.AuthorityMenu, err error) {
	// 查询的是视图表authority_menu
	err = global.SYS_DB.Where("authority_id = ?", info.AuthorityId).Order("sort").Find(&menus).Error
	return
}

// AddMenuAuthority 增加 1角色->N菜单 的关联关系
func (m *MenuService) AddMenuAuthority(menus []model.BaseMenu, authorityId string) error {
	var auth model.Authority
	auth.AuthorityId = authorityId
	auth.BaseMenus = menus
	err := AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

// GetInfoList 获取菜单列表，包括子菜单
func (m *MenuService) GetInfoList() (list interface{}, total int64, err error) {
	var menuList []model.BaseMenu
	treemap, err := getBaseMenuTreeMap()
	menuList = treemap["0"] // 顶级menu
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treemap)
	}
	return menuList, total, err
}

// getBaseMenuTreeMap 获取路由总树map,格式为 map[父级菜单ID][]Menu
func getBaseMenuTreeMap() (treeMap map[string][]model.BaseMenu, err error) {
	var allMenus []model.BaseMenu
	treeMap = make(map[string][]model.BaseMenu)
	err = global.SYS_DB.Order("sort").Find(&allMenus).Error // 排序查询
	for _, v := range allMenus {
		// map每个元素是：相同parent_id的所有menu切片. 即：同级别menu在同一个k的slice中
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// getBaseChildrenList 递归获取子菜单
func getBaseChildrenList(menu *model.BaseMenu, treeMap map[string][]model.BaseMenu) (err error) {
	// 菜单都已经查询并处理放在map了，只需要从map中拿即可
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
