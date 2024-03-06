package logger

import (
	"bytes"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	SdkLogger      = logrus.New()
	CmbcLogger     = logrus.New()
	AlibabaLogger  = logrus.New()
	LianlianLogger = logrus.New()
	KuaijieLoger   = logrus.New()
	ErrorLogger    = logrus.New()
	HeliLogger     = logrus.New()
	RdsLogger      *RdsLoggerWriter
)

type CustomLogger struct {
	File   string
	Logger *logrus.Logger
}

func New(LogFilePath string, RunMode string) {
	var err error
	_, err = os.Stat(LogFilePath)
	switch {
	case os.IsNotExist(err):
		err = os.MkdirAll(LogFilePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	case os.IsPermission(err):
		panic(err)
	}
	for _, item := range []CustomLogger{
		{File: "error.log", Logger: ErrorLogger},
		{File: "sdk.log", Logger: SdkLogger},
		{File: "cmbc.log", Logger: CmbcLogger},
		{File: "alibaba.log", Logger: AlibabaLogger},
		{File: "lianlian.log", Logger: LianlianLogger},
		{File: "kuaijie.log", Logger: KuaijieLoger},
		{File: "heli.log", Logger: HeliLogger},
	} {
		loggerToFile(LogFilePath, item.File, item.Logger, RunMode)
	}
	RdsLogger = newRdsLoggerWriter(LogFilePath, "rds.log", logrus.New(), RunMode)
}

type LogFormatter struct{}

func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	var file string
	var l int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		l = entry.Caller.Line
	}
	//fmt.Println(entry.Data)
	msg := fmt.Sprintf(
		"[GOID:%d] [%s] %s [%s:%d] %s\n",
		getGID(),
		strings.ToUpper(entry.Level.String()),
		timestamp, file, l,
		entry.Message)
	return []byte(msg), nil
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func loggerToFile(LogFilePath string, LogFileName string, LoggerInstance *logrus.Logger, RunMode string) {

	fileName := path.Join(LogFilePath, LogFileName)

	//writer := bufio.NewWriter(src)
	//LoggerInstance.SetOutput(writer)
	outSelect(LoggerInstance, RunMode)
	LoggerInstance.SetReportCaller(true)
	//LoggerInstance.SetFormatter(new(LogFormatter))

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

	lfHook := lfshook.NewHook(writeMap, &LogFormatter{})

	// 新增 Hook
	LoggerInstance.AddHook(lfHook)
}

//
//func accessLoggerToFile(LogFilePath string, LogFileName string, LoggerInstance *logrus.Logger, RunMode string) {
//
//	fileName := path.Join(LogFilePath, LogFileName)
//
//	//writer := bufio.NewWriter(src)
//	//LoggerInstance.SetOutput(writer)
//	outSelect(LoggerInstance, RunMode)
//	LoggerInstance.SetReportCaller(true)
//	//LoggerInstance.SetFormatter(new(LogFormatter))
//
//	logWriter, _ := rotatelogs.New(
//		// 分割后的文件名称
//		fileName+".%Y%m%d.log",
//
//		// 生成软链，指向最新日志文件
//		rotatelogs.WithLinkName(fileName),
//
//		// 设置最大保存时间(7天)
//		rotatelogs.WithMaxAge(30*24*time.Hour),
//
//		// 设置日志切割时间间隔(1天)
//		rotatelogs.WithRotationTime(24*time.Hour),
//	)
//
//	writeMap := lfshook.WriterMap{
//		logrus.InfoLevel:  logWriter,
//		logrus.FatalLevel: logWriter,
//		logrus.DebugLevel: logWriter,
//		logrus.WarnLevel:  logWriter,
//		logrus.ErrorLevel: logWriter,
//		logrus.PanicLevel: logWriter,
//	}
//
//	lfHook := lfshook.NewHook(writeMap, &LogFormatter{})
//
//	// 新增 Hook
//	LoggerInstance.AddHook(lfHook)
//}
//
//func routerLoggerToFile(LogFilePath string, LogFileName string, LoggerInstance *logrus.Logger, RunMode string) {
//
//	fileName := path.Join(LogFilePath, LogFileName)
//	//src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//writer := bufio.NewWriter(src)
//	//LoggerInstance.SetOutput(src)
//	outSelect(LoggerInstance, RunMode)
//
//	logWriter, _ := rotatelogs.New(
//		// 分割后的文件名称
//		fileName+".%Y%m%d.log",
//
//		// 生成软链，指向最新日志文件
//		rotatelogs.WithLinkName(fileName),
//
//		// 设置最大保存时间(7天)
//		rotatelogs.WithMaxAge(30*24*time.Hour),
//
//		// 设置日志切割时间间隔(1天)
//		rotatelogs.WithRotationTime(24*time.Hour),
//	)
//
//	writeMap := lfshook.WriterMap{
//		logrus.InfoLevel:  logWriter,
//		logrus.FatalLevel: logWriter,
//		logrus.DebugLevel: logWriter,
//		logrus.WarnLevel:  logWriter,
//		logrus.ErrorLevel: logWriter,
//		logrus.PanicLevel: logWriter,
//	}
//
//	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{})
//
//	// 新增 Hook
//	LoggerInstance.AddHook(lfHook)
//}

func outSelect(LoggerInstance *logrus.Logger, RunMode string) {

	if RunMode == "release" {
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		LoggerInstance.SetOutput(src)
	} else {
		LoggerInstance.SetOutput(os.Stdout)
	}
}
