package config

type Zap struct {
	Level string `mapstructure:"level" json:"level" yaml:"level"` // 级别
	Dir   string `mapstructure:"dir" json:"dir"  yaml:"dir"`      // 日志文件夹
}
