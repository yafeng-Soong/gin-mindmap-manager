package config

type Server struct {
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
}
