package tmpl

// ConfigContent 配置文件
var ConfigContent = `// 配置文件导入yaml文件是configstruct.go
//
// 配置文件可以使用 -c 的参数
// https://github.com/go-yaml/yaml
package config

import (
	"path"
	"runtime"
	"fmt"

	"github.com/spf13/viper"
)

// 设置配置文件的 环境变量
var (
	//MysqlDbName 数据库名称
	MysqlDbName string
	// MysqlPassword 数据库密码
	MysqlPassword string
	// MysqlUsername 连接数据库用户名
	MysqlUsername string
	// MysqlPort 数据库端口号
	MysqlPort string
	// MysqlHost 数据库主机
	MysqlHost string
	// MysqlConnect gorm连接数据库信息
	MysqlConnect string
	// LogDirector 日志目录
	LogDirector string
	// LogInfoFile info日志文件
	LogInfoFilename string
	LogMaxSize int
	LogMaxBackups int
	LogMaxAge int
	LogLevel string
)

// 获取文件绝对路径
func getCurrPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func InitConfig(opt *Option) {
	viper.SetConfigFile(opt.ConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	MysqlDbName = viper.GetString("db.mysql.dbname")
	MysqlPassword = viper.GetString("Db.Mysql.Password")
	MysqlUsername = viper.GetString("Db.Mysql.Username")
	MysqlPort = viper.GetString("Db.Mysql.Port")
	MysqlHost = viper.GetString("Db.Mysql.Host")

	// "user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlConnect = MysqlUsername + ":" + MysqlPassword + "@(" + MysqlHost + ":" + MysqlPort + ")/" + MysqlDbName +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	LogDirector = viper.GetString("Log.LogDirector")
	if LogDirector == ""{
		LogDirector = path.Join(path.Dir(getCurrPath()), "log")
	}
	LogInfoFilename = path.Join(LogDirector, viper.GetString("log.logInfoFilename"))
	LogMaxSize = viper.GetInt("log.logMaxSize")
	LogMaxBackups = viper.GetInt("log.LogMaxBackups")
	LogMaxAge = viper.GetInt("log.LogMaxAge")
	LogLevel = viper.GetString("logLevel")
}

`
