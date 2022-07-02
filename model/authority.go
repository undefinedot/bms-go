package model

import (
	"time"
)

// Authority 角色表
type Authority struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time  `sql:"index"`
	AuthorityId   string      `json:"authority_id,omitempty"`
	AuthorityName string      `json:"authority_name,omitempty"`
	ParentId      string      `json:"parent_id,omitempty"`
	Children      []Authority `json:"children,omitempty"`
	Users         []User      `json:"users,omitempty"`
	DefaultRouter string      `json:"default_router,omitempty"`
}

func (Authority) TableName() string {
	return "authorities"
}
