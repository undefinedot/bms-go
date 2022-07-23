package request

// GetById 按id查找user
type GetById struct {
	ID int `json:"id" binding:"required"`
}

// GetAuthorityId 按id查找authority
type GetAuthorityId struct {
	AuthorityId string `json:"authority_id" binding:"required"`
}

// PageInfo 分页
type PageInfo struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	KeyWord  string `json:"key_word"`
}
