package config

type System struct {
	Name          string `mapstructure:"name" json:"name" yaml:"name"`                              // 项目名称
	Version       string `mapstructure:"version" json:"version" yaml:"version"`                     // 项目版本
	Port          string `mapstructure:"port" json:"port" yaml:"port"`                              // 端口值
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截 限制ip等
}
