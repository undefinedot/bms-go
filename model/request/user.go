package request

// Register 用户注册
type Register struct {
	NiceName     string   `json:"nice_name" binding:"required"`
	Username     string   `json:"username" binding:"required"`
	Password     string   `json:"password" binding:"required"`
	RePassword   string   `json:"re_password" binding:"eqfield=Password"`
	HeaderImg    string   `json:"header_img"`
	Phone        string   `json:"phone"`
	Email        string   `json:"email"`
	AuthorityId  string   `json:"authority_id"`
	AuthorityIds []string `json:"authority_ids"`
}

// Login 用户登录
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	// todo: 增加验证码功能
}

// ChangePwd 修改密码
type ChangePwd struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

// ChangeUserInfo 修改用户信息
type ChangeUserInfo struct {
	ID           uint     `json:"id" binding:"required"`
	NickName     string   `json:"nick_name"`
	HeaderImg    string   `json:"header_img"`
	Phone        string   `json:"phone"`
	Email        string   `json:"email"`
	AuthorityIds []string `json:"authority_ids"`
}
