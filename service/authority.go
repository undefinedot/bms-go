package service

import (
	"bms-go/global"
	"bms-go/model"
	"bms-go/model/common/request"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AuthorityService struct{}

func (as *AuthorityService) CreateAuthority(auth model.Authority) (model.Authority, error) {
	var authRet model.Authority
	err := global.SYS_DB.Where("authority_id = ?", auth.AuthorityId).First(&authRet).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return authRet, errors.New("角色已存在！")
	}
	err = global.SYS_DB.Create(&auth).Error
	return auth, err
}

func (as *AuthorityService) UpdateAuthority(auth model.Authority) (model.Authority, error) {
	err := global.SYS_DB.Updates(&auth).Error
	return auth, err
}
func (as *AuthorityService) DeleteAuthority(auth model.Authority) error {
	if errors.Is(global.SYS_DB.Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("角色不存在")
	}
	if len(auth.Users) > 0 {
		return errors.New("当前角色下还有用户，禁止删除！")
	}
	// user表中有authority_id字段
	if !errors.Is(global.SYS_DB.Where("authority_id = ?", auth.AuthorityId).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("还有用户属于当前角色，禁止删除！")
	}
	// 角色下的子角色
	if !errors.Is(global.SYS_DB.Where("parent_id = ?", auth.AuthorityId).First(&model.Authority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("当前角色下还有子角色，禁止删除！")
	}
	// 硬删除角色
	if err := global.SYS_DB.Unscoped().Delete(&auth).Error; err != nil {
		return err
	}
	// 删除中间表user_authority,user表中在前面已判断
	err := global.SYS_DB.Where("authority_id = ?", auth.AuthorityId).Delete(&model.UserAuthority{}).Error
	return err
}

func (a *AuthorityService) GetAuthorityInfoList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	offset := (pageInfo.Page - 1) * pageInfo.PageSize
	limit := pageInfo.PageSize
	var authorities []model.Authority
	db := global.SYS_DB.Model(&model.Authority{})
	err = db.Where("parent_id = ?", "0").Count(&total).Error
	// 查出所有顶级角色
	err = db.Offset(offset).Limit(limit).Where("parent_id = ?", "0").Find(&authorities).Error
	// 递归查询
	if len(authorities) > 0 {
		for k := range authorities {
			err = findChildrenAuthority(&authorities[k])
		}
	}
	return authorities, total, err
}

// findChildrenAuthority 递归查出所有子角色
func findChildrenAuthority(auth *model.Authority) error {
	err := global.SYS_DB.Where("parent_id = ?", auth.AuthorityId).Find(&auth.Children).Error
	if len(auth.Children) > 0 {
		for k := range auth.Children {
			err = findChildrenAuthority(&auth.Children[k])
		}
	}
	return err
}
