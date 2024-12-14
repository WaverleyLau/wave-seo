package config

type ServerConf struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
