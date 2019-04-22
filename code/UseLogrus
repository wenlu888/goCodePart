package tools

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"path"
	"time"
)

// 你可以创建很多instance
// var log = logrus.New()

// 日志本地文件分
func NewHook(logName string ,logPath string ,maxRemainCnt uint) log.Hook {

	baseLogPath := path.Join(logPath,logName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H",
		// WithLinkName 为最新的日志建立软连接，以方便随时找到当前日志文件
		// rotatelogs.WithLinkName(logName),

		// WithRotationTime 设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Minute),

		// WithMaxAge和WithRotationCount 二者只能设置一个，
		// WithMaxAge 设置文件清理前的最长保存时间，
		// WithRotationCount 设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}


	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})

	return lfsHook
}

func MRecover() {
	if err := recover(); err != nil {
		var logger = log.WithFields(log.Fields{
			"err": err,
		})
		logger.Error("recover from :", err)
		// debug.PrintStack()
	}
}


func CheckErr(err error, desc string) {
	if err != nil {
		logger := log.WithFields(log.Fields{
			"err":     string(err.Error()),
		})
		logger.Error(desc)
	}
}
