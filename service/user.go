package service

import (
	"bms-go/global"
	"bms-go/model"
	"bms-go/model/common/request"
	"bms-go/utils"

	uuid "github.com/satori/go.uuid"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserService struct{}

// todo: 用户昵称

func (us *UserService) Register(u *model.User) (*model.User, error) {
	var user model.User
	// 用户是否已存在
	if !errors.Is(global.SYS_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名已注册")
	}
	// 生成uuid，加密
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	// 注册
	err := global.SYS_DB.Create(&u).Error
	return u, err
}

func (us *UserService) Login(u *model.User) (*model.User, error) {
	var user model.User
	err := global.SYS_DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		// 验证密码
		if ok := utils.CheckPassword(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		// todo: 对应的角色，权限相关
	}
	return &user, err
}

func (us *UserService) ChangePassword(u *model.User, newPwd string) (*model.User, error) {
	var user model.User
	err := global.SYS_DB.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return nil, err
	}
	if ok := utils.CheckPassword(newPwd, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPwd)
	err = global.SYS_DB.Save(&user).Error
	return &user, err
}

func (us *UserService) SetUserInfo(u model.User) error {
	return global.SYS_DB.Updates(&u).Error
}

func (us *UserService) DeleteUser(uid int) error {
	var user model.User
	return global.SYS_DB.Where("id = ?", uid).Delete(&user).Error
}

func (us *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	offset := (info.Page - 1) * info.PageSize
	limit := info.PageSize
	db := global.SYS_DB.Model(&model.User{})
	var userList []model.User
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return userList, total, err
}

func (us *UserService) GetUserInfo(uuid uuid.UUID) (*model.User, error) {
	var user model.User
	err := global.SYS_DB.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
