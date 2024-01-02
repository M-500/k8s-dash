package cfg

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/2 17:44
//

type RedisCfg struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Pass string `mapstructure:"pass" json:"pass"`
}
