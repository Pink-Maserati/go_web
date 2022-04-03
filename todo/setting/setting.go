package setting

import "gopkg.in/ini.v1"

var Conf = new(AppConfig)

//AppConfig 应用程序配置
type AppConfig struct {
	Release       bool `ini:"release"`
	Port          int  `ini:"port"`
	*MySQLConfig  `ini:"mysql"`
	*LoggerConfig `ini:"logger"`
}

//MySQLConfig 配置
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	DB       string `ini:"db"`
}

type LoggerConfig struct {
	Path string `ini:"path"`
	Name string `ini:"name"`
}

//初始化配置
func Init(file string) error {
	return ini.MapTo(Conf, file)

}
