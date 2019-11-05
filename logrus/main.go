package main

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	logg "github.com/sirupsen/logrus"
	"os"
	"time"
)

// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。
var log = logg.New()

func newLfsHook(logLevel *string, maxRemainCnt uint) logg.Hook {

	logName := "test"
	writer, err := rotatelogs.New(
		logName+".%Y%m%d%H",
		// WithLinkName 为最新的日志建立软连接，以方便随时找到当前日志文件
		rotatelogs.WithLinkName(logName),

		// WithRotationTime 设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour),

		// WithMaxAge和WithRotationCount 二者只能设置一个，
		// WithMaxAge 设置文件清理前的最长保存时间，
		// WithRotationCount 设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		logg.Errorf("config local file system for logger error: %v", err)
	}

	logg.SetLevel(logg.WarnLevel)

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logg.DebugLevel: writer,
		logg.InfoLevel:  writer,
		logg.WarnLevel:  writer,
		logg.ErrorLevel: writer,
		logg.FatalLevel: writer,
		logg.PanicLevel: writer,
	}, &logg.TextFormatter{DisableColors: true})

	return lfsHook
}

func main() {
	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	log.Out = os.Stdout
	log.AddHook(newLfsHook(nil, 2))
	// 为当前logrus实例设置消息输出格式为json格式。
	// 同样地，也可以单独为某个logrus实例设置日志级别和hook，这里不详细叙述。
	log.Formatter = &logg.JSONFormatter{PrettyPrint: true}

	log.WithFields(logg.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
	log.Info("oon")
}
