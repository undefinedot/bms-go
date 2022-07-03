package model

type UserAuthority struct {
	UserID      uint   `gorm:"column:user_id"`
	AuthorityID string `gorm:"column:authority_id"`
}

func (*UserAuthority) TableName() string {
	return "user_authority"
}
