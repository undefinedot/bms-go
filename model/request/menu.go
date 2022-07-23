package request

import "bms-go/model"

type AddMenuAuthorityInfo struct {
	Menus       []model.BaseMenu `json:"menus"`
	AuthorityId string           `json:"authority_id"`
}
