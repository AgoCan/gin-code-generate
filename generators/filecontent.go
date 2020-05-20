package generators

// 配置内容
const (
	// MainContent main.go 内容
	MainContent = `package main

func main(){
}
`
	// ConfigContent config/config.go 内容
	ConfigContent = `// 配置文件导入yaml文件是configstruct.go
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

	ConfigYamlContent = `# 使用yaml做配置项
# 数据库配置项
db:
  mysql:
    dbname: "example"
    password: "root1234"
    username: "root"
    port: 3306
    host: "127.0.0.1"

log:
  # 默认路径是运行程序的目录
#  logDirector: ./logging
  logAutoFile: log.middleware
  logInfoFile: log.manual
`

	ConfigTestContent = `package config

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestConfig(t *testing.T){
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		t.Error("so such file or director")
	}
	fmt.Println(string(data))
	fileData := Conf.getConfig()
	fmt.Printf("%T \n%v\n" ,fileData,fileData)
}
`


	ConfigStructContent = `
package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

// Config 配置项结构体
type Config struct {
	Db struct {
		Mysql struct {
			DbName   string ` + "`yaml:\"dbname\"`" + `
			Password string ` + "`yaml:\"password\"`" + `
			Username string ` + "`yaml:\"username\"`" + `
			Port     string ` + "`yaml:\"port\"`" + `
			Host     string ` + "`yaml:\"host\"`" + `
		}
	}
	Log struct{
		LogDirector string	` + "`yaml:\"logDirector\"`" + `
		LogAutoFile string ` + "`yaml:\"logAutoFile\"`" + `
		//logWaringFile string
		LogInfoFile string ` + "`yaml:\"logInfoFile\"`" + `
	}
}

// Conf 配置项
var Conf Config

// 获取文件绝对路径
func getCurrPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

var yamlFilePath string

// 获取环境变量
func (c *Config) getConfig() *Config {
	yamlFilePath = getCurrPath()
	configYaml, err := ioutil.ReadFile(yamlFilePath + "/config.yaml")
	if err != nil {
		fmt.Printf("err %v\n", err)
	}
	err = yaml.Unmarshal(configYaml, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}`


	// MiddlewareLogContent 导入包的时候 %s 拼接项目名称
	MiddlewareLogContent = `package middleware

/*
# 日志文件默认存储位置
 '/var/logging/das-go'
使用配置文件进行配置

https://github.com/sirupsen/logrus

*/

import (
	"%s/utils/logging"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	" %s/config"
)

// SetInitLog 初始化
func LogMiddleware() gin.HandlerFunc {
	// 日志对应yaml配置文件logAutoFile
	logging.Init(config.LogAutoFile)
	return func (c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		Method := c.Request.Method

		// 请求路由
		RequestURI := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logging.Logger.WithFields(logrus.Fields{
			"status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"method"   : Method,
			"request_uri"      : RequestURI,
		}).Info()

	}

}


`
	RouterContent = `package routers

import (
	"github.com/gin-gonic/gin"

	"%s/middleware"

	"%s/handler/docker"
)

// SetupRouter 路由路口
func SetupRouter() *gin.Engine{
	router := gin.Default()
	// docker 相关操作
	router.Use(middleware.LogMiddleware())
	return router
}`
)

func init(){

}
