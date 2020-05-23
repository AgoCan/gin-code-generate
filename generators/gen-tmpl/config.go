package tmpl

// ConfigContent 配置文件
var ConfigContent = `// 配置文件导入yaml文件是configstruct.go
//
// 配置文件可以使用 -c 的参数
// https://github.com/go-yaml/yaml
package config

import (
	"path"
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
	LogAutoFile string
	// LogWaringFile waring 日志文件
	//LogWaringFile string
	//// LogErrorFile  error 日志文件
	LogInfoFile string
)



func init() {
	Conf.getConfig()
	MysqlDbName = Conf.Db.Mysql.DbName
	MysqlPassword = Conf.Db.Mysql.Password
	MysqlUsername = Conf.Db.Mysql.Username
	MysqlPort = Conf.Db.Mysql.Port
	MysqlHost = Conf.Db.Mysql.Host
	// "user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlConnect = MysqlUsername + ":" + MysqlPassword + "@(" + MysqlHost + ":" + MysqlPort + ")/" + MysqlDbName +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	LogDirector = Conf.Log.LogDirector
	if LogDirector == ""{
		LogDirector = path.Join(path.Dir(getCurrPath()), "log")
	}
	LogAutoFile = path.Join(LogDirector, Conf.Log.LogAutoFile)
	//LogWaringFile := path.Join(LogDirector, Conf.logging.logWaringFile)
	LogInfoFile = path.Join(LogDirector, Conf.Log.LogInfoFile)

}

`
