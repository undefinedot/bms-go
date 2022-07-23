package model

// AuthorityMenu 视图
type AuthorityMenu struct {
	BaseMenu
	MenuId      string
	AuthorityId string
	Children    []AuthorityMenu
}

func (AuthorityMenu) TableName() string {
	return "authority_menu"
}
