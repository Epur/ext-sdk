package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

//定义自己的Writer
type RdsLoggerWriter struct {
	logger *logrus.Logger
}

//实现gorm/logger.Writer接口
func (r *RdsLoggerWriter) Printf(format string, v ...interface{}) {
	r.logger.Info(fmt.Sprintf(format, v...))
}

type RdsLogFormatter struct{}

func (s *RdsLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	//var file string
	//var l int
	//if entry.Caller != nil {
	//	file = filepath.Base(entry.Caller.File)
	//	l = entry.Caller.Line
	//}
	//fmt.Println(entry.Data)
	msg := fmt.Sprintf("%s %s\n", timestamp, entry.Message)
	return []byte(msg), nil
}

func newRdsLoggerWriter(LogFilePath string, LogFileName string, LoggerInstance *logrus.Logger, RunMode string) *RdsLoggerWriter {

	//配置logrus

	fileName := path.Join(LogFilePath, LogFileName)

	outSelect(LoggerInstance, RunMode)

	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(30*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &RdsLogFormatter{})

	// 新增 Hook
	LoggerInstance.AddHook(lfHook)

	return &RdsLoggerWriter{logger: LoggerInstance}
}
