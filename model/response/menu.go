package response

import "bms-go/model"

// BaseMenuRes 单条查询
type BaseMenuRes struct {
	Menu model.BaseMenu `json:"menu"`
}

// BaseMenusRes 多条查询
type BaseMenusRes struct {
	Menus []model.BaseMenu `json:"menus"`
}
