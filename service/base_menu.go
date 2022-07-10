package service

import (
	"bms-go/global"
	"bms-go/model"
	"errors"

	"gorm.io/gorm"
)

type BaseMenuService struct{}

// AddBaseMenu 新增菜单
func (b *BaseMenuService) AddBaseMenu(menu model.BaseMenu) error {
	if !errors.Is(global.SYS_DB.Where("name = ?", menu.Name).First(&model.BaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("菜单名称已存在！")
	}
	return global.SYS_DB.Create(&menu).Error
}

func (b *BaseMenuService) DeleteBaseMenu(id int) error {
	// 菜单有子菜单时禁止删除
	err := global.SYS_DB.Where("parent_id = ?", id).First(&model.BaseMenu{}).Error
	if err != nil {
		// 删除menu和中间表数据
		var menu model.BaseMenu
		err = global.SYS_DB.Where("id = ?", id).First(&menu).Select("Authorities").Delete(&menu).Error
	} else {
		return errors.New("菜单有子菜单，禁止删除！")
	}
	return err
}

func (b *BaseMenuService) UpdateBaseMenu(menu model.BaseMenu) (err error) {
	var oldMenu model.BaseMenu
	db := global.SYS_DB.Where("id = ?", menu.ID).Find(&oldMenu)
	// 是否已存在
	if oldMenu.Name != menu.Name {
		// 修改后的name不能与已存在的name相同
		if !errors.Is(global.SYS_DB.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&model.BaseMenu{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同name")
		}
	}
	return db.Updates(&menu).Error // 使用map可以保存零值
}

func (b *BaseMenuService) GetBaseMenuById(id int) (menu model.BaseMenu, err error) {
	err = global.SYS_DB.Where("id = ?", id).First(&menu).Error
	return
}
