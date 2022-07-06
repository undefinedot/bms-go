package model

type UserAuthority struct {
	UserID      uint   `gorm:"column:user_id"`
	AuthorityID string `gorm:"column:authority_authority_id"` // 易错：gorm生成的中间表字段跟想的不一样
}

func (*UserAuthority) TableName() string {
	return "user_authority"
}
