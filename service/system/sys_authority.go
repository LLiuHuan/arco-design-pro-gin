// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 17:54
package system

import (
	"errors"
	"strconv"

	"github.com/lliuhuan/arco-design-pro-gin/global"

	"github.com/lliuhuan/arco-design-pro-gin/model/common/request"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/lliuhuan/arco-design-pro-gin/model/system/response"
	"gorm.io/gorm"
)

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

//findChildrenAuthority 查询子角色
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.SysAuthority
//@return: err error
func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.AdpDb.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

//CreateAuthority 创建一个角色
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority
func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (err error, authority system.SysAuthority) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.AdpDb.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = global.AdpDb.Create(&auth).Error
	return err, auth
}

//CopyAuthority 复制一个角色
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: err error, authority model.SysAuthority
func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority system.SysAuthority) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.AdpDb.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []system.SysAuthority{}
	err, menus := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	if err != nil {
		return
	}
	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.AdpDb.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}
	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = CasbinServiceApp.UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}

//UpdateAuthority 更改一个角色
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority
func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority) (err error, authority system.SysAuthority) {
	err = global.AdpDb.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Updates(&auth).Error
	return err, auth
}

//DeleteAuthority 删除角色
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *model.SysAuthority
//@return: err error
func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) (err error) {
	if !errors.Is(global.AdpDb.Where("authority_id = ?", auth.AuthorityId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.AdpDb.Where("parent_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.AdpDb.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	if len(auth.SysBaseMenus) > 0 {
		err = global.AdpDb.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		if err != nil {
			return
		}
		//err = db.Association("SysBaseMenus").Delete(&auth)
	} else {
		err = db.Error
		if err != nil {
			return
		}
	}
	err = global.AdpDb.Delete(&[]system.SysUseAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error
	CasbinServiceApp.ClearCasbin(0, auth.AuthorityId)
	return err
}

//GetAuthorityInfoList 分页获取数据
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64
func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.AdpDb.Model(&system.SysAuthority{})
	err = db.Where("parent_id = ?", "0").Count(&total).Error
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = authorityService.findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

//GetAuthorityInfo 获取所有角色信息
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.SysAuthority
//@return: err error, sa model.SysAuthority
func (authorityService *AuthorityService) GetAuthorityInfo(auth system.SysAuthority) (err error, sa system.SysAuthority) {
	err = global.AdpDb.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

//SetDataAuthority 设置角色资源权限
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: SetDataAuthority
//@description: 设置角色资源权限
//@param: auth model.SysAuthority
//@return: error
func (authorityService *AuthorityService) SetDataAuthority(auth system.SysAuthority) error {
	var s system.SysAuthority
	global.AdpDb.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.AdpDb.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

//SetMenuAuthority 菜单与角色绑定
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error
func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	var s system.SysAuthority
	global.AdpDb.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.AdpDb.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}
