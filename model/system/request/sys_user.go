// Package request
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 17:34
package request

// Login User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// SetUserAuth Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}

// SetUserAuthorities Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}

// Register User register structure
type Register struct {
	Username     string   `json:"userName"`
	Password     string   `json:"passWord"`
	NickName     string   `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg    string   `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	AuthorityId  string   `json:"authorityId" gorm:"default:888"`
	AuthorityIds []string `json:"authorityIds"`
}
