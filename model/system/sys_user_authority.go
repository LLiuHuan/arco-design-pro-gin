// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 17:49
package system

type SysUseAuthority struct {
	SysUserId               uint   `gorm:"column:sys_user_id"`
	SysAuthorityAuthorityId string `gorm:"column:sys_authority_authority_id"`
}

func (s *SysUseAuthority) TableName() string {
	return "sys_user_authority"
}
