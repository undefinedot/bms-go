package model

import (
	"bms-go/global"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	global.BaseModel
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户uuid"`
	Username  string    `json:"username" gorm:"comment:用户登录名"`
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	Password  string    `json:"-"`
	HeaderImg string    `json:"headerImg" gorm:"default:https://i1.hdslb.com/bfs/face/379d097c2ea8c9be2c83819d326e6ad1ebe8c811.jpg;comment:头像"`
	//AuthorityId string
	//Authority   Authority
	//Authorities []Authority
	Phone       string      `json:"phone" gorm:"comment:手机号码"`
	Email       string      `json:"email" gorm:"邮箱"`
	AuthorityID string      `gorm:"column:authority_id;default:888;comment:对应的角色id"`
	Authorities []Authority `gorm:"many2many:user_authority"`
}

func (User) TableName() string {
	return "users"
}
