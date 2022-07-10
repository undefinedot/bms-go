package model

import "bms-go/global"

type BaseMenu struct {
	global.BaseModel
	MenuLevel   uint        `json:"-"`
	ParentId    string      `json:"parentId" gorm:"comment:父菜单ID"`                    // 父菜单ID
	Path        string      `json:"path" binding:"required" gorm:"comment:对应的路由path"` // 对应的路由path
	Name        string      `json:"name" binding:"required" gorm:"comment:路由name"`    // 路由name
	Hidden      bool        `json:"hidden" gorm:"comment:是否在列表隐藏"`                    // 是否在列表隐藏
	Sort        int         `json:"sort" gorm:"comment:排序标记"`                         // 排序标记
	Authorities []Authority `json:"authorities" gorm:"many2many:authority_menus"`
	Children    []BaseMenu  `json:"children" gorm:"-"`
}

func (BaseMenu) TableName() string {
	return "base_menus"
}
