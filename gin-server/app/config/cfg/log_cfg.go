package cfg

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/2 17:44
//

type LogCfg struct {
	Level         string `mapstructure:"level" json:"level"`
	Path          string `mapstructure:"path" json:"path"`
	Filename      string `mapstructure:"filename" json:"filename"`
	MaxSize       int    `mapstructure:"max_size" json:"maxSize"`
	MaxAge        int    `mapstructure:"max_age" json:"maxAge"`
	MaxBackups    int    `mapstructure:"max_backups" json:"maxBackups"`
	Compress      bool   `mapstructure:"compress" json:"compress"`
	StacktraceKey string `mapstructure:"stacktrace_key" json:"stacktraceKey"`
	Format        string `mapstructure:"format" json:"format"`
}
