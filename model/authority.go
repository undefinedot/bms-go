package model

import (
	"time"
)

// Authority 角色表
type Authority struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time  `gorm:"index" json:"-"`
	AuthorityId   string      `gorm:"primaryKey;comment:角色id;size:90" json:"authority_id" binding:"required"`
	AuthorityName string      `json:"authority_name" binding:"required"`
	ParentId      string      `gorm:"column:parent_id;comment:角色名称" json:"parent_id" binding:"required"`
	Children      []Authority `json:"children"`
	Users         []User      `gorm:"many2many:user_authority" json:"-"`
	BaseMenus     []BaseMenu  `gorm:"many2many:authority_menus" json:"menus"`
}

func (Authority) TableName() string {
	return "authorities"
}
