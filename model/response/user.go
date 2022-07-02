package response

import "bms-go/model"

type UserRes struct {
	User *model.User `json:"user,omitempty"`
}

type LoginRes struct {
	User   model.User `json:"user"`
	Token  string     `json:"token"`
	Expire int64      `json:"expire"`
}

// PageResult 分页用户列表
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}
