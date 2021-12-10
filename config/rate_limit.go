// Package config
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 14:41
package config

type RateLimit struct {
	IpVerify   bool   `mapstructure:"ip-verify" json:"ip-verify" yaml:"ip-verify"`          // 是否打开ip限流
	IpLimitCon int64  `mapstructure:"ip-limit-con" json:"ip-limit-con" yaml:"ip-limit-con"` // 每秒访问ip超过多少次
	IpListKey  string `mapstructure:"ip-list-key" json:"ip-list-key" yaml:"ip-list-key"`    // ip列表的key

	Cap     int64 `mapstructure:"cap" json:"cap" yaml:"cap"`             // 初始化数量
	Quantum int64 `mapstructure:"quantum" json:"quantum" yaml:"quantum"` // 每秒增加数量

	LimitCountIP int `mapstructure:"iplimit-count" json:"iplimitCount" yaml:"iplimit-count"` // IP限制次数 LimitTimeIP 内 限制 LimitCountIP 次
	LimitTimeIP  int `mapstructure:"iplimit-time" json:"iplimitTime" yaml:"iplimit-time"`    // IP限制一个小时 单位秒 3600
}
