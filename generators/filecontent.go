package generators

// 配置内容
const (
	// MainContent main.go 内容
	MainContent = `package main

import (
	"fmt"

	"%s/model"
	"%s/routers"

)

var (
	err error
)

func main() {
	// 连接数据库并在代码结束后关闭
	err = model.InitMysql()
	if err != nil {
		// 数据库连接失败，直接报错
		panic(err)
	}
	defer model.Close()

	// 调用路由组
	router := routers.SetupRouter()

	err = router.Run(":9000")
	if err != nil {
		fmt.Println(err)
	}
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

	"%s/config"
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

)

// SetupRouter 路由路口
func SetupRouter() *gin.Engine{
	router := gin.Default()
	router.Use(middleware.LogMiddleware())
	return router
}`
	ModelContent = `package model
// https://gorm.io/zh_CN/
import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"%s/config"
)

// DB db handler
var DB *gorm.DB

// InitMysql 初始化数据库
func InitMysql() (err error) {
	DB, err = gorm.Open("mysql", config.MysqlConnect)
	if err != nil {
		return err
	}
	// 创建数据库
	migrate()
	// 尝试连接
	err = DB.DB().Ping()
	return
}

// Close 关闭数据库
func Close()(){

	err := DB.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func migrate(){

}
`
	UtilsLogContent = `package logging

import (
	"%s/config"
	"fmt"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// 自定义日志内容

var Logger *logrus.Logger

// 初始化日志的钩子
func Init(LogInfoFile string){
	// 日志path
	logFilePath := LogInfoFile
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		// panic("no such file or director")
	}
	Logger = logrus.New()
	Logger.Out = file
	// 设置日志级别
	Logger.SetLevel(logrus.InfoLevel)

	// 设置 rotateLogs
	logWriter, err := rotateLogs.New(
		// 分割后的文件名称
		logFilePath + ".%%Y%%m%%d",

		// 生成软链，指向最新日志文件
		rotateLogs.WithLinkName(logFilePath),

		// 设置最大保存时间(7天)
		rotateLogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotateLogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic("rotate logging faild")
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	// 使用 json 记录数据
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	Logger.AddHook(lfHook)
}

// 定义日志级别

const (
	InfoLevel = "info"
	WarnLevel = "warning"
	ErrorLevel = "ErrorLevel"
	PanicLevel = "PanicLevel"
)

// AppendLog 自定义添加日志内容
func Append(logType string, content interface{}){
	// 必须要初始化Logger
	Init(config.LogInfoFile)
	// 当前时间
	now := time.Now()
	// 根据打印的级别进行分类
	switch logType {
	case "info":
		Logger.WithFields(logrus.Fields{
			"time": now,
		}).Info(content)
	case "warning":
		Logger.WithFields(logrus.Fields{
			"time": now,
			"content": content,
		}).Warn()
	case "error":
		Logger.WithFields(logrus.Fields{
			"time": now,
			"content": content,
		}).Error()
	case "panic":
		Logger.WithFields(logrus.Fields{
			"time": now,
			"content": content,
		}).Panic()
	default:
		Logger.WithFields(logrus.Fields{
			"time": now,
			"content": content,
		}).Info()
	}

}`

	GoModContent = `module %s

go 1.14
require (
	github.com/containerd/containerd v1.3.3 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v0.0.0-20191113042239-ea84732a7725
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/godoctor/godoctor v0.0.0-20181123222458-69df17f3a6f6 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/josharian/impl v0.0.0-20191119165012-6b9658ad00c7 // indirect
	github.com/mdempsky/gocode v0.0.0-20191202075140-939b4a677f2f // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518 // indirect
	github.com/willf/bitset v1.1.10 // indirect
	golang.org/x/tools v0.0.0-20200312045724-11d5b4c81c7d // indirect
	google.golang.org/grpc v1.28.0 // indirect
)

`
)

