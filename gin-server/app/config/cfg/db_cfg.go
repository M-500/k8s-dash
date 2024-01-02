package cfg

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/2 17:44
//

type MysqlCfg struct {
	Engine          string `mapstructure:"engine"       json:"engine"`             // mysql
	Username        string `mapstructure:"username"      json:"username"`          // root
	Password        string `mapstructure:"password"      json:"password"`          // 123456
	Host            string `mapstructure:"host"          json:"host"`              // 127.0.0.1
	Port            int    `mapstructure:"port"          json:"port"`              // 3306
	Database        string `mapstructure:"database"      json:"database"`          // user_growth
	Charset         string `mapstructure:"charset"       json:"charset"`           // utf8mb4
	ShowSQL         bool   `mapstructure:"showSQL"       json:"showSQL"`           // false
	MaxIdleConns    int    `mapstructure:"maxIdleConns"  json:"maxIdleConns"`      // 2
	MaxOpenConns    int    `mapstructure:"maxOpenConns"  json:"maxOpenConns"`      // 10
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime" json:"connMaxLifetime"` // 30m
}
