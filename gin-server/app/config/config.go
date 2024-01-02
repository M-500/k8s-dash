package config

import (
	"gin-server/app/config/cfg"
	"github.com/spf13/viper"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/2 17:44
//

var ConfigInstance *Config

type Config struct {
	Name  string       `mapstructure:"name"`
	Host  string       `mapstructure:"host"`
	Local string       `mapstructure:"local"`
	Port  int          `mapstructure:"port"`
	DB    cfg.MysqlCfg `mapstructure:"db"`
	Redis cfg.RedisCfg `mapstructure:"redis"`
}

func MustLoadCfg(path string, types string) *Config {
	if types == "YML" {
		return LoadYmlConfig(path)
	}
	return LoadYmlConfig(path)
}
func LoadYmlConfig(path string) *Config {
	pc := &Config{}
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(pc); err != nil {
		panic(err)
	}
	ConfigInstance = pc
	return ConfigInstance
}
