package system

import (
	"bms-go/global"
	"bms-go/model"
	"bms-go/model/common/request"
	"bms-go/model/common/response"
	reqModel "bms-go/model/request"
	resModel "bms-go/model/response"
	"bms-go/utils"
	"fmt"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (b *BaseApi) Login(ctx *gin.Context) {
	var r reqModel.Login
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	// todo: 验证码
	user := &model.User{
		Username: r.Username,
		Password: r.Password,
	}
	userRet, err := userService.Login(user)
	if err != nil {
		zap.L().Error("userService.Login failed", zap.Error(err))
		response.FailWithMsg("用户名或密码错误", ctx)
		return
	}
	// 签发jwt
	b.TokenNext(ctx, userRet)
}

func (b *BaseApi) Register(ctx *gin.Context) {
	var r reqModel.Register
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	// 处理角色id集合，string=>[]Authority
	var authorities []model.Authority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, model.Authority{
			AuthorityId: v,
		})
	}
	user := &model.User{
		Username:    r.Username,
		NickName:    r.NiceName,
		Password:    r.Password,
		HeaderImg:   r.HeaderImg,
		Phone:       r.Phone,
		Email:       r.Email,
		AuthorityID: r.AuthorityId,
		Authorities: authorities,
	}
	userRet, err := userService.Register(user)
	if err != nil {
		zap.L().Error("userService.Register failed", zap.Error(err))
		response.FailWithMsg("用户注册失败", ctx)
		return
	}
	zap.L().Info("用户注册成功", zap.Uint("userID", userRet.ID))
	response.OkWithDetail(resModel.UserRes{User: userRet}, "注册成功", ctx)
}

// TokenNext 登录成功后签发jwt
func (b *BaseApi) TokenNext(ctx *gin.Context, user *model.User) {
	j := utils.NewJWT()
	claims := j.CreateClaims(reqModel.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		Username: user.Username,
		NickName: user.NickName,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		zap.L().Error("TokenNext j.CreateClaims 生成token失败", zap.Error(err))
		response.FailWithMsg("获取token失败", ctx)
		return
	}
	// todo: 多点登录, redis
	// redis: 根据username在redis取jwt,没取到才设置新jwt
	// 若已存在jwt,拉黑旧jwt,设置新jwt
	response.OkWithDetail(resModel.LoginRes{
		User:   *user,
		Token:  token,
		Expire: claims.StandardClaims.ExpiresAt,
	}, "登录成功", ctx)
}

func (b *BaseApi) ChangePassword(ctx *gin.Context) {
	var user reqModel.ChangePwd
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	u := &model.User{Username: user.Username, Password: user.Password}
	if _, err := userService.ChangePassword(u, user.NewPassword); err != nil {
		zap.L().Error("修改密码失败", zap.Error(err))
		response.FailWithMsg("修改密码失败，原密码错误", ctx)
		return
	}
	response.OkWithMsg("密码修改成功", ctx)
}

func (b *BaseApi) SetUserInfo(ctx *gin.Context) {
	var user reqModel.ChangeUserInfo
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	// 更新用户的角色列表
	if len(user.AuthorityIds) != 0 {
		err := userService.SetUserAuthorities(user.ID, user.AuthorityIds)
		if err != nil {
			zap.L().Error("SetUserAuthorities() failed", zap.Error(err))
			response.FailWithMsg("设置用户信息失败！", ctx)
			return
		}
	}
	// 更新用户表
	u := model.User{
		BaseModel: global.BaseModel{ID: user.ID},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
	}
	if err := userService.SetUserInfo(u); err != nil {
		zap.L().Error("userService.SetUserInfo failed", zap.Error(err))
		response.FailWithMsg("设置失败", ctx)
		return
	}
	response.OkWithMsg("设置成功", ctx)
}

func (b *BaseApi) DeleteUser(ctx *gin.Context) {
	var reqId request.GetById
	if err := ctx.ShouldBindJSON(&reqId); err != nil {
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	// 禁止删除自己
	jwtId := utils.GetUserID(ctx) // todo: 带jwt测试删除用户
	if jwtId == uint(reqId.ID) {
		response.FailWithMsg("禁止删除自己", ctx)
		return
	}

	if err := userService.DeleteUser(reqId.ID); err != nil {
		zap.L().Error("userService.DeleteUser failed", zap.Error(err))
		response.FailWithMsg("删除失败", ctx)
		return
	}
	response.FailWithMsg("删除成功", ctx)
}

// GetUserList 用户列表分页
func (b *BaseApi) GetUserList(ctx *gin.Context) {
	var pageInfo = request.PageInfo{
		// 分页默认值,第一页page=1
		Page:     1,
		PageSize: 3,
	}
	if err := ctx.ShouldBindJSON(&pageInfo); err != nil {
		fmt.Println(err)
		response.FailWithMsg(err.Error(), ctx)
		return
	}
	if list, total, err := userService.GetUserInfoList(pageInfo); err != nil {
		zap.L().Error("分页获取失败", zap.Error(err))
		response.FailWithMsg("分页获取用户列表失败!", ctx)
		return
	} else {
		response.OkWithDetail(resModel.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

func (b *BaseApi) GetUserInfo(ctx *gin.Context) {
	uuid := utils.GetUserUuid(ctx)
	if user, err := userService.GetUserInfo(uuid); err != nil {
		zap.L().Error("GetUserInfo by uuid failed", zap.Error(err))
		response.FailWithMsg("获取失败", ctx)
	} else {
		response.OkWithDetail(gin.H{"userinfo": user}, "获取成功", ctx)
	}
}
