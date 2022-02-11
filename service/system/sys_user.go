package system

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/common/request"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/lliuhuan/arco-design-pro-gin/utils"
)

type UserService struct {
}

// Login 用户登陆
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser
func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	if nil == global.AdpDb {
		return fmt.Errorf("db not init"), nil
	}

	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.AdpDb.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	return err, &user
}

//Register 用户注册
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser
func (userService *UserService) Register(u system.SysUser) (err error, userInter system.SysUser) {
	var user system.SysUser
	if !errors.Is(global.AdpDb.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.AdpDb.Create(&u).Error
	return err, u
}

//ChangePassword 修改用户密码
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (err error, userInter *system.SysUser) {
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.AdpDb.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//GetUserInfoList 分页查询用户数据
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64
func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.AdpDb.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return err, userList, total
}

//SetUserAuthority 设置一个用户的权限
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error
func (userService *UserService) SetUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	assignErr := global.AdpDb.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUseAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.AdpDb.Where("uuid = ?", uuid).First(&system.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//SetUserAuthorities 设置一个用户的权限
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error
func (userService *UserService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.AdpDb.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUseAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysUseAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//DeleteUser 删除用户
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error
func (userService *UserService) DeleteUser(id float64) (err error) {
	var user system.SysUser
	err = global.AdpDb.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.AdpDb.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
	return err
}

//SetUserInfo 设置用户信息
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser
func (userService *UserService) SetUserInfo(reqUser system.SysUser) (err error, user system.SysUser) {
	err = global.AdpDb.Updates(&reqUser).Error
	return err, reqUser
}

//GetUserInfo 获取用户信息
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (err error, user system.SysUser) {
	var reqUser system.SysUser
	err = global.AdpDb.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	var menu MenuService
	err1, menus := menu.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: reqUser.AuthorityId})
	if err1 == nil {
		for _, sysMenu := range menus {
			if sysMenu.MenuType == 3 && sysMenu.Meta.Permissions != "" {
				reqUser.Permissions = append(reqUser.Permissions, sysMenu.Meta.Permissions)
			}
		}
	}
	return err, reqUser
}

//FindUserById 通过id获取用户信息
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser
func (userService *UserService) FindUserById(id int) (err error, user *system.SysUser) {
	var u system.SysUser
	err = global.AdpDb.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//FindUserByUuid 通过uuid获取用户信息
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser
func (userService *UserService) FindUserByUuid(uuid string) (err error, user *system.SysUser) {
	var u system.SysUser
	if err = global.AdpDb.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

//ResetPassword 修改用户密码
//@author: [lliuhuan](https://github.com/lliuhuan)
//@function: resetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error
func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.AdpDb.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.MD5V([]byte("123456"))).Error
	return err
}
