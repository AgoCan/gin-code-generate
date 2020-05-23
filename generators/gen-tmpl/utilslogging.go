package tmpl

// UtilsLoggingContent 日志工具
var UtilsLoggingContent = `package logging
import (
	"{{ .ProjectName }}/config"
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
